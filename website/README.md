# Lume Documentation Website

A multilingual static documentation website for Lume with language switching support.

## Supported Languages

- 🇺🇸 **English** (en)
- 🇨🇳 **中文** (zh)
- 🇪🇸 **Español** (es)

## Features

- 🌍 **Language Switcher** - Dropdown menu in the navbar
- 💾 **Persistence** - Language preference saved to localStorage
- ⌨️ **Keyboard Shortcut** - `Ctrl/Cmd + Shift + L` to cycle languages
- 📋 **Copy to Clipboard** - One-click copy for install command
- 🎨 **Dark Theme** - Modern dark UI design
- 📱 **Responsive** - Mobile-friendly layout

## Project Structure

```
website/
├── index.html          # Main page
├── css/
│   └── style.css       # Stylesheet
├── js/
│   └── i18n.js         # Internationalization logic
├── content/
│   ├── en.json         # English translations
│   ├── zh.json         # Chinese translations
│   └── es.json         # Spanish translations
└── README.md           # This file
```

## Local Development

### Option 1: Python HTTP Server

```bash
cd website
python3 -m http.server 8080
```

Then open http://localhost:8080

### Option 2: Node.js (http-server)

```bash
npm install -g http-server
cd website
http-server -p 8080
```

### Option 3: VS Code Live Server

Install the "Live Server" extension and right-click on `index.html` → "Open with Live Server"

## Deployment

### GitHub Pages

1. Go to repository Settings → Pages
2. Source: Deploy from a branch
3. Branch: `main` / `root` or `main` / `docs`
4. If using `docs` folder, move `website/` contents to `docs/`
5. Save and wait for deployment

### Custom Domain

1. Add a `CNAME` file in the website root with your domain
2. Configure DNS to point to GitHub Pages
3. Enable HTTPS in repository settings

## Adding a New Language

1. Copy `content/en.json` to `content/{lang}.json`
2. Translate all values in the JSON file
3. Add language info to `js/i18n.js`:
   ```javascript
   langNames: {
       // ... existing languages
       {lang}: { flag: '🇯🇵', name: 'JA', full: '日本語' }
   }
   ```
4. Add language option to `index.html` dropdown
5. Update `html.lang` mapping in `js/i18n.js` if needed

## Browser Support

- Chrome 80+
- Firefox 75+
- Safari 13.1+
- Edge 80+

## License

MIT License - Same as the Lume project
