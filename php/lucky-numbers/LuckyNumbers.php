<?php

class LuckyNumbers
{
    public function sumUp(array $digitsOfNumber1, array $digitsOfNumber2): int
    {
        return (int) implode($digitsOfNumber1) + (int) implode($digitsOfNumber2);
    }

    public function isPalindrome(int $number): bool
    {
        $forward = (string) $number;
        $backward = implode(array_reverse(str_split($forward)));
        // $backward = strrev($forward); // 🤡

        return $forward == $backward;
    }

    public function validate(string $input): string
    {
        if ($input == '') {
            return "Required field";
        } else if ((int) $input <= 0) {
            return "Must be a whole number larger than 0";
        }

        return '';
    }
}
