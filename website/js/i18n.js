// ===== i18n Configuration =====
const i18n = {
    currentLang: localStorage.getItem('lume-lang') || 'en',
    
    // Language display names
    langNames: {
        en: { flag: '🇺🇸', name: 'EN', full: 'English' },
        zh: { flag: '🇨🇳', name: 'ZH', full: '中文' },
        es: { flag: '🇪🇸', name: 'ES', full: 'Español' }
    },

    // Translations
    translations: {}
};

// ===== Load Translations =====
async function loadTranslations(lang) {
    if (i18n.translations[lang]) {
        return i18n.translations[lang];
    }
    
    try {
        const response = await fetch(`content/${lang}.json`);
        const data = await response.json();
        i18n.translations[lang] = data;
        return data;
    } catch (error) {
        console.error(`Failed to load translations for ${lang}:`, error);
        return null;
    }
}

// ===== Apply Translations =====
function applyTranslations(lang, translations) {
    if (!translations) return;
    
    // Find all elements with data-i18n attribute
    const elements = document.querySelectorAll('[data-i18n]');
    
    elements.forEach(el => {
        const key = el.getAttribute('data-i18n');
        const value = getNestedValue(translations, key);
        
        if (value) {
            // Check if content contains HTML
            if (value.includes('<') && value.includes('>')) {
                el.innerHTML = value;
            } else {
                el.textContent = value;
            }
        }
    });
    
    // Update HTML lang attribute
    document.documentElement.lang = lang === 'zh' ? 'zh-CN' : lang;
    
    // Update language button
    updateLangButton(lang);
    
    // Update active state in dropdown
    updateActiveLangOption(lang);
}

// ===== Get Nested Object Value =====
function getNestedValue(obj, path) {
    return path.split('.').reduce((current, key) => {
        return current && current[key] !== undefined ? current[key] : null;
    }, obj);
}

// ===== Update Language Button =====
function updateLangButton(lang) {
    const btn = document.getElementById('langBtn');
    const langInfo = i18n.langNames[lang];
    if (btn && langInfo) {
        const currentSpan = btn.querySelector('.lang-current');
        if (currentSpan) {
            currentSpan.textContent = `${langInfo.flag} ${langInfo.name}`;
        }
    }
}

// ===== Update Active Language Option =====
function updateActiveLangOption(lang) {
    const options = document.querySelectorAll('.lang-option');
    options.forEach(option => {
        option.classList.remove('active');
        if (option.getAttribute('data-lang') === lang) {
            option.classList.add('active');
        }
    });
}

// ===== Set Language =====
async function setLanguage(lang) {
    if (!i18n.langNames[lang]) return;
    
    i18n.currentLang = lang;
    localStorage.setItem('lume-lang', lang);
    
    const translations = await loadTranslations(lang);
    applyTranslations(lang, translations);
}

// ===== Initialize =====
async function initI18n() {
    // Load initial language
    await setLanguage(i18n.currentLang);
    
    // Setup language switcher
    setupLangSwitcher();
    
    // Setup copy buttons
    setupCopyButtons();
}

// ===== Setup Language Switcher =====
function setupLangSwitcher() {
    const langBtn = document.getElementById('langBtn');
    const langDropdown = document.getElementById('langDropdown');
    const langSwitcher = document.querySelector('.lang-switcher');
    
    if (!langBtn || !langDropdown) return;
    
    // Toggle dropdown
    langBtn.addEventListener('click', (e) => {
        e.stopPropagation();
        langSwitcher.classList.toggle('active');
    });
    
    // Close dropdown when clicking outside
    document.addEventListener('click', () => {
        langSwitcher.classList.remove('active');
    });
    
    langDropdown.addEventListener('click', (e) => {
        e.stopPropagation();
    });
    
    // Language option click
    const langOptions = langDropdown.querySelectorAll('.lang-option');
    langOptions.forEach(option => {
        option.addEventListener('click', (e) => {
            e.preventDefault();
            const lang = option.getAttribute('data-lang');
            setLanguage(lang);
            langSwitcher.classList.remove('active');
        });
    });
    
    // Keyboard shortcut for language switcher
    document.addEventListener('keydown', (e) => {
        // Ctrl/Cmd + Shift + L to toggle language
        if ((e.ctrlKey || e.metaKey) && e.shiftKey && e.key === 'L') {
            e.preventDefault();
            const langs = Object.keys(i18n.langNames);
            const currentIndex = langs.indexOf(i18n.currentLang);
            const nextIndex = (currentIndex + 1) % langs.length;
            setLanguage(langs[nextIndex]);
        }
    });
}

// ===== Setup Copy Buttons =====
function setupCopyButtons() {
    const copyBtns = document.querySelectorAll('.copy-btn');
    
    copyBtns.forEach(btn => {
        btn.addEventListener('click', async () => {
            const text = btn.getAttribute('data-clipboard');
            if (!text) return;
            
            try {
                await navigator.clipboard.writeText(text);
                
                // Show copied state
                btn.classList.add('copied');
                const originalHTML = btn.innerHTML;
                btn.innerHTML = `
                    <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                        <polyline points="20 6 9 17 4 12"></polyline>
                    </svg>
                `;
                
                setTimeout(() => {
                    btn.classList.remove('copied');
                    btn.innerHTML = originalHTML;
                }, 2000);
            } catch (err) {
                console.error('Failed to copy:', err);
            }
        });
    });
}

// ===== Start =====
document.addEventListener('DOMContentLoaded', initI18n);
