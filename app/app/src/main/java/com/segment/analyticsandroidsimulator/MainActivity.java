package com.segment.analyticsandroidsimulator;

import android.app.Activity;
import android.os.Bundle;
import com.segment.analytics.Analytics;

public class MainActivity extends Activity {
  @Override protected void onCreate(Bundle savedInstanceState) {
    super.onCreate(savedInstanceState);
    setContentView(R.layout.activity_main);

    Analytics.with(this).track("Main Activity Started");
  }
}
