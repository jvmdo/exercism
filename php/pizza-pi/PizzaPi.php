<?php

class PizzaPi
{
    public function calculateDoughRequirement(int $number_of_pizzas, int $number_of_persons): int
    {
        return $number_of_pizzas * (($number_of_persons * 20) + 200);
    }

    public function calculateSauceRequirement(int $number_of_pizzas, int $sauce_can_volume, float $sauce_per_pizza = 125.0): int
    {
        return intval($number_of_pizzas * $sauce_per_pizza / $sauce_can_volume);
    }

    public function calculateCheeseCubeCoverage(int $cheese_dimension, float $thickness, float $pizza_diameter): int
    {
        return floor(($cheese_dimension ** 3) / ($thickness * pi() * $pizza_diameter));
    }

    public function calculateLeftOverSlices(int $number_of_pizzas, int $number_of_friends): int
    {
        return (8 * $number_of_pizzas) % $number_of_friends;
    }
}
