# Dragon Vision Prediction

import numpy as np

def predict_future_battles():
    # Simulate dragon vision prediction for future battles
    clouds = ["Nimbus", "Cumulus", "Stratus"]
    predictions = np.random.choice(clouds, size=5, replace=True)
    return predictions

# Perform dragon vision prediction
future_battles = predict_future_battles()
print("Dragon Vision Predictions for Future Battles:")
for i, prediction in enumerate(future_battles):
    print(f"{i+1}. {prediction}")
