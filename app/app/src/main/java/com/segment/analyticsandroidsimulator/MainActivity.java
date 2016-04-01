package com.segment.analyticsandroidsimulator;

import android.app.Activity;
import android.content.Intent;
import android.os.Bundle;
import android.util.Log;
import com.segment.analytics.Analytics;
import timber.log.Timber;

public class MainActivity extends Activity {
  @Override protected void onCreate(Bundle savedInstanceState) {
    super.onCreate(savedInstanceState);

    setContentView(R.layout.activity_main);
    handleIntent(getIntent());
  }

  @Override protected void onNewIntent(Intent intent) {
    super.onNewIntent(intent);

    handleIntent(intent);
  }

  void handleIntent(Intent intent) {
    if (intent == null) {
      Timber.d("Main activity launched with no intent.");
      return;
    }
    String type = intent.getStringExtra("type");
    if (type == null) {
      Timber.d("No type provided in intent.");
      return;
    }

    switch (type) {
      case "track":
        track(intent);
        break;
      case "flush":
        flush();
        break;
      default:
        throw new IllegalArgumentException("invalid event type: " + type);
    }
  }

  void track(Intent intent) {
    String event = intent.getStringExtra("event");
    Analytics.with(this).track(event);
  }

  void flush() {
    Analytics.with(this).flush();
  }
}
