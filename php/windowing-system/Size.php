<?php

class Size
{
  public int $width, $height;

  function __construct(int $height, int $width)
  {
    $this->height = $height;
    $this->width = $width;
  }
}
