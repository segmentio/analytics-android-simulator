package com.segment.analyticsandroidsimulator;

final class Utils {
  private Utils() {
    throw new AssertionError("no instances");
  }

  static boolean stringEqual(String s1, String s2) {
    return s1.compareToIgnoreCase(s2) == 0;
  }
}
