<?php

class Position
{
  public int $x, $y;

  function __construct(int $y, int $x)
  {
    $this->y = $y;
    $this->x = $x;
  }
}
