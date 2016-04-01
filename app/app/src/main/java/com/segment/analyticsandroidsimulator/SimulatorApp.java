package com.segment.analyticsandroidsimulator;

import android.app.Application;
import com.segment.analytics.Analytics;
import timber.log.Timber;

public class SimulatorApp extends Application {

  @Override public void onCreate() {
    super.onCreate();

    Timber.plant(new Timber.DebugTree());

    Analytics analytics = new Analytics.Builder(this, "4txshy8l73") //
        .logLevel(Analytics.LogLevel.VERBOSE) //
        .build();
    Analytics.setSingletonInstance(analytics);
  }
}
