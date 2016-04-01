package com.segment.analyticsandroidsimulator;

import android.app.Activity;
import android.content.Intent;
import android.os.Bundle;
import com.segment.analytics.Analytics;
import com.segment.analytics.Properties;
import java.io.IOException;
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
        try {
          track(intent);
        } catch (IOException e) {
          Timber.e(e, "Error running track.");
        }
        break;
      case "flush":
        flush();
        break;
      default:
        throw new IllegalArgumentException("Invalid event type: " + type);
    }
  }

  void track(Intent intent) throws IOException {
    String event = intent.getStringExtra("event");
    Bundle extras = intent.getExtras();

    Properties properties = new Properties();
    for (String key : extras.keySet()) {
      if (key.startsWith("properties_")) {
        properties.put(key.replaceAll("properties_", ""), extras.get(key));
      }
    }

    Timber.d("analytics.track(%s, %s);", event, properties);
    Analytics.with(this).track(event, properties);
  }

  void flush() {
    Timber.d("analytics.flush();");
    Analytics.with(this).flush();
  }
}
