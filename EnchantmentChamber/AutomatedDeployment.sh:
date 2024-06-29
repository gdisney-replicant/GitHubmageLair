// Build Enchantments

const { exec } = require('child_process');

// Enchantment for compiling JavaScript
function compileJS() {
    console.log('Compiling JavaScript files...');
    exec('tsc', (error, stdout, stderr) => {
        if (error) {
            console.error(`Error: ${error.message}`);
            return;
        }
        if (stderr) {
            console.error(`stderr: ${stderr}`);
            return;
        }
        console.log(`Compilation successful: ${stdout}`);
    });
}

// Enchantment for optimizing CSS
function optimizeCSS() {
    console.log('Optimizing CSS files...');
    exec('postcss styles.css -o styles.min.css', (error, stdout, stderr) => {
        if (error) {
            console.error(`Error: ${error.message}`);
            return;
        }
        if (stderr) {
            console.error(`stderr: ${stderr}`);
            return;
        }
        console.log(`Optimization successful: ${stdout}`);
    });
}

// Enchantment for bundling assets
function bundleAssets() {
    console.log('Bundling assets...');
    exec('webpack', (error, stdout, stderr) => {
        if (error) {
            console.error(`Error: ${error.message}`);
            return;
        }
        if (stderr) {
            console.error(`stderr: ${stderr}`);
            return;
        }
        console.log(`Bundling successful: ${stdout}`);
    });
}

// Apply build enchantments
compileJS();
optimizeCSS();
bundleAssets();
