# Magical Utilities

class Spell:
    def __init__(self, name, description):
        self.name = name
        self.description = description

    def cast(self):
        print(f"Casting {self.name} spell: {self.description}")

class GitHubmage:
    def __init__(self, name):
        self.name = name
        self.spells = []

    def learn_spell(self, spell):
        self.spells.append(spell)

    def cast_spells(self):
        print(f"{self.name} casts the following spells:")
        for spell in self.spells:
            spell.cast()

# Define spells
fireball_spell = Spell("Fireball", "Summons a fiery projectile to incinerate enemies")
heal_spell = Spell("Heal", "Restores health to the GitHubmage")
teleport_spell = Spell("Teleport", "Instantly transports the GitHubmage to a chosen location")

# Create GitHubmage character
githubmage = GitHubmage("Merlin")

# Learn spells
githubmage.learn_spell(fireball_spell)
githubmage.learn_spell(heal_spell)
githubmage.learn_spell(teleport_spell)

# Cast spells
githubmage.cast_spells()
