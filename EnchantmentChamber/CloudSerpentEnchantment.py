# Cloud Serpent Enchantment

class CloudSerpent:
    def __init__(self, name, color):
        self.name = name
        self.color = color
        self.length = 10  # Default length

    def grow(self, growth_factor):
        self.length *= growth_factor

    def fly(self):
        print(f"The {self.color} cloud serpent gracefully takes flight, soaring through the skies.")

    def breathe_fire(self):
        print("The cloud serpent inhales deeply, exhaling a stream of fiery clouds.")

# Create a majestic cloud serpent
azure_serpent = CloudSerpent("Azure Serpent", "blue")
azure_serpent.grow(2)  # Double the size of the serpent
azure_serpent.fly()  # Command the serpent to fly
azure_serpent.breathe_fire()  # Command the serpent to breathe fire
