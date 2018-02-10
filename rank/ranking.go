package rank

func Calculate(
	ranks *Rank,
	algorithm func(int, int, int, int, int, int, int, int, int) float32,
) {
	updateRanks(ranks, algorithm)
}

func updateRanks(
	ranks *Rank,
	algorithm func(int, int, int, int, int, int, int, int, int) float32,
) {
	for x, xMap := range ranks.Relation.Node {
		for y, _ := range xMap {
			qty := ranks.Relation.Node[x][y].Qty

			if ranks.Relation.Max < qty {
				ranks.Relation.Max = qty
			}

			if ranks.Relation.Min > qty || ranks.Relation.Min == 0 {
				ranks.Relation.Min = qty
			}
		}
	}

	for _, word := range ranks.Words {
		if ranks.Max < word.Qty {
			ranks.Max = word.Qty
		}

		if ranks.Min > word.Qty || ranks.Min == 0 {
			ranks.Min = word.Qty
		}
	}

	for _, word := range ranks.Words {
		weight := algorithm(word.ID, 0, 0, 0, 0, word.Qty, 0, ranks.Min, ranks.Max)
		word.Weight = weight
	}

	for x, xMap := range ranks.Relation.Node {
		for y, _ := range xMap {
			qty := ranks.Relation.Node[x][y].Qty
			sentenceIDs := ranks.Relation.Node[x][y].SentenceIDs
			weight := algorithm(
				x,
				y,
				qty,
				ranks.Relation.Min,
				ranks.Relation.Max,
				ranks.Words[x].Qty,
				ranks.Words[y].Qty,
				ranks.Min,
				ranks.Max,
			)
			ranks.Relation.Node[x][y] = Score{ranks.Relation.Node[x][y].Qty, weight, sentenceIDs}
		}
	}
}