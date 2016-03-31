package com.segment.analyticsandroidsimulator;

import android.app.Application;
import com.segment.analytics.Analytics;

public class SimulatorApp extends Application {

  @Override public void onCreate() {
    super.onCreate();

    Analytics analytics = new Analytics.Builder(this, "4txshy8l73") //
        .logLevel(Analytics.LogLevel.VERBOSE) //
        .build();
    Analytics.setSingletonInstance(analytics);
  }
}
