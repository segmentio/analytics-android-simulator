package com.segment.analyticsandroidsimulator;

import android.app.Application;
import com.segment.analytics.Analytics;

public class SimulatorApp extends Application {

  @Override public void onCreate() {
    super.onCreate();

    Analytics.setSingletonInstance(new Analytics.Builder(this, "4txshy8l73").build());
  }
}
