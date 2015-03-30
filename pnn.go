package main

type PNN struct {
    // public double cases[][];
    cases [][]float64
    // public double controls[][];
    controls [][]float64
    // public double inputs[][];
    inputs [][]float64
    // public ArrayList<Double> classifications;
    classifications []float64
    // public double correctClassifications[];
    correctClassifications []float64
    // public File inputFile;
    // public File patternFile;
    // public double casePrior;
    casePrior float64
    // public double contPrior;
    contPrior float64
    // public double caseLoss;
    caseLoss float64
    // public double contLoss;
    contPrior float64
    // public ArrayList<Double> sigma;
    sigma []float64
    // public int columns = 412;
    columns int
    // public int count = 0;
    count int
    // public int accum = 0;
    accum int
    // public double bestAccuracy;
    bestAccuracy float64
    // public double bestSigma;
    bestSigma float64
}

func NewPNN() *PNN {
    pnn := PNN{}
    return &pnn
}
