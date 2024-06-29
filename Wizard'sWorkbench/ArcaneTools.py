# Arcane Tools

class SpellBook:
    def __init__(self, name):
        self.name = name
        self.spells = []

    def learn_spell(self, spell):
        self.spells.append(spell)
        print(f"Learned new spell: {spell}")

    def cast_spell(self, spell_name):
        if spell_name in self.spells:
            print(f"Casting {spell_name}...")
            # Code to cast spell goes here
            print("Spell cast successfully!")
        else:
            print(f"Spell '{spell_name}' not found in {self.name}.")

# Create a SpellBook instance
my_spellbook = SpellBook("Book of Spells")

# Learn new spells
my_spellbook.learn_spell("Fireball")
my_spellbook.learn_spell("Teleportation")

# Cast a spell
my_spellbook.cast_spell("Fireball")
