package com.segment.analyticsandroidsimulator;

import android.app.Activity;
import android.content.Intent;
import android.os.Bundle;
import android.util.Log;
import com.segment.analytics.Analytics;

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
      Log.d("Simulator", "No intent.");
      return;
    }
    String type = intent.getStringExtra("type");
    if (type == null) {
      Log.d("Simulator", "No type.");
      return;
    }

    switch (type) {
      case "track":
        Log.d("Simulator", "Track.");
        track(intent);
        break;
      case "flush":
        Log.d("Simulator", "Flush.");
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
