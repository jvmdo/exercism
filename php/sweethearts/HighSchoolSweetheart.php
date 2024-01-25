<?php

class HighSchoolSweetheart
{
    public function firstLetter(string $name): string
    {
        return trim($name)[0];
    }

    public function initial(string $name): string
    {
        return strtoupper($this->firstLetter($name)) . '.';
    }

    public function initials(string $name): string
    {
        [$first_name, $last_name] = explode(" ", $name);

        $first_initial = $this->initial($first_name);
        $last_initial = $this->initial($last_name);

        return "$first_initial $last_initial";
    }

    public function pair(string $sweetheart_a, string $sweetheart_b): string
    {
        $first_sweetheart = $this->initials($sweetheart_a);
        $second_sweetheart = $this->initials($sweetheart_b);

        return <<<HEART
              ******       ******
            **      **   **      **
          **         ** **         **
         **            *            **
         **                         **
         **     $first_sweetheart  +  $second_sweetheart     **
          **                       **
            **                   **
              **               **
                **           **
                  **       **
                    **   **
                      ***
                       *
         HEART;
    }
}
