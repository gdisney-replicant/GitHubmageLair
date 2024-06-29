// Dragon Scale Enhancement

class DragonScale {
    constructor(color, hardness) {
        this.color = color;
        this.hardness = hardness;
        this.weight = 10;  // Default weight in kilograms
    }

    increaseWeight(weightIncrease) {
        this.weight += weightIncrease;
    }

    enchant() {
        console.log(`The ${this.color} dragon scale gleams with magical energy, enhancing its ${this.hardness} hardness.`);
    }

    shapeshift() {
        console.log(`The ${this.color} dragon scale begins to shimmer, shifting its form into a new shape.`);
    }
}

// Create a dragon scale and apply enhancements
const rubyScale = new DragonScale("ruby", "diamond-like");
rubyScale.increaseWeight(5);  // Increase the weight of the scale
rubyScale.enchant();  // Apply enchantment to the scale
rubyScale.shapeshift();  // Command the scale to shapeshift
