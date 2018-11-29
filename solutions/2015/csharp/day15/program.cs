using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;

namespace Day15
{
    public static class Program {
		public struct Ingredient {
			public string Name;
			public int Capacity;
			public int Durability;
			public int Flavor;
			public int Texture;
			public int Calories;
		}

		static Ingredient ParseEntry(string input) {
			var inputParts = input.Split(' ');
			var i = new Ingredient();
			i.Name = inputParts[0].Replace(",","");
			i.Capacity = int.Parse(inputParts[2].Replace(",",""));
			i.Durability = int.Parse(inputParts[4].Replace(",",""));
			i.Flavor = int.Parse(inputParts[6].Replace(",",""));
			i.Texture = int.Parse(inputParts[8].Replace(",",""));
			i.Calories = int.Parse(inputParts[10].Replace(",",""));
			return i;
		}

		static void ReadFileByLines(string fileName, Action<String> lineAction) {
			using(var fs = System.IO.File.Open(fileName, System.IO.FileMode.Open)) {
				using(var sr = new System.IO.StreamReader(fs)) {
					while(!sr.EndOfStream) {
						var line = sr.ReadLine();
						lineAction(line);
					}
				}
			}
		}

		static int CalculateScore(List<int> distribution, List<Ingredient> ingredients) {
			var capacity = 0;
			var durability = 0;
			var flavor = 0;
			var texture = 0;
			var calories = 0;

			for(var x = 0; x < distribution.Count; x++) {
				var amount = distribution[x];
				var i = ingredients[x];

				capacity = capacity + (amount * i.Capacity);
				durability = durability + (amount * i.Durability);
				flavor = flavor + (amount * i.Flavor);
				texture = texture + (amount * i.Texture);
				calories = calories + (amount * i.Calories);
			}

			var subTotal = (capacity > 0 ? capacity : 0) *
				(durability > 0 ? durability : 0) *
				(flavor > 0 ? flavor : 0) *
				(texture > 0 ? texture : 0);
			return subTotal;
		}

		static List<List<int>> PermuteDistributions(int total, int buckets) {
			return PermuteDistributionsFromExisting(total, buckets, new List<int>());
		}

		static List<List<int>> PermuteDistributionsFromExisting(int total, int buckets, List<int> existing) {
			var output = new List<List<int>>();
			var existingLength = existing.Count;
			var existingSum = sum(existing);
			var remainder = total - existingSum;

			if (buckets == 1) {
				var newExisting = new List<int>(existing);
				newExisting.Add(remainder);
				output.Add(newExisting);
				return output;
			}

			for (var x = 0; x <= remainder; x++) {
				var newExisting = new List<int>(existing);
				newExisting.Add(x);

				var results = PermuteDistributionsFromExisting(total, buckets-1, newExisting);
				foreach(var result in results) {
					output.Add(result);
				}
			}

			return output;
		}

		static int sum(List<int> values) {
			var total = 0;
			for (var x = 0; x < values.Count; x++) {
				total += values[x];
			}

			return total;
		}

		static void Main(string[] args) {
			var codeFile = "../../testdata/day15";
			var ingredients = new List<Ingredient>();
			ReadFileByLines(codeFile, (line) => {
				ingredients.Add(ParseEntry(line));
			});

			var distributions = PermuteDistributions(100, ingredients.Count);
			var bestScore = 0;
			//var bestDistribution = new List<int>();
			foreach(var distribution in distributions) {
				var score = CalculateScore(distribution, ingredients);
				if (score > bestScore) {
					bestScore = score;
					//bestDistribution = distribution.ToList();
				}
			}

			Console.WriteLine("Best Score: " + bestScore.ToString());
		}
	}
}