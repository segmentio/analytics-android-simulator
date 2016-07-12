package com.segment.analytics.android.cli;

import android.content.Intent;
import android.os.Bundle;
import android.support.v7.app.AppCompatActivity;
import com.segment.analytics.Analytics;
import com.segment.analytics.Properties;
import com.segment.analytics.Traits;
import timber.log.Timber;

public class MainActivity extends AppCompatActivity {
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
    Timber.plant(new Timber.DebugTree());
    if (intent == null) {
      Timber.d("Main activity launched with no intent.");
      return;
    }
    String type = intent.getStringExtra("type");
    if (type == null) {
      Timber.d("No type provided in intent.");
      return;
    }

    // TODO: writeKey as flag.
    // https://segment.com/segment-engineering/sources/android-test/settings/keys
    Analytics analytics = new Analytics.Builder(this, "5m6gbdgho6") //
        .logLevel(Analytics.LogLevel.VERBOSE) //
        .build();
    Analytics.setSingletonInstance(analytics);

    switch (type) {
      case "track":
        track(intent);
        break;
      case "screen":
        screen(intent);
        break;
      case "identify":
        identify(intent);
        break;
      case "alias":
        alias(intent);
        break;
      case "group":
        group(intent);
        break;
      case "flush":
        flush();
        break;
      case "reset":
        reset();
        break;
      default:
        throw new IllegalArgumentException("Invalid event type: " + type);
    }
  }

  void track(Intent intent) {
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

  void screen(Intent intent) {
    String name = intent.getStringExtra("name");
    String category = intent.getStringExtra("category");

    Bundle extras = intent.getExtras();
    Properties properties = new Properties();
    for (String key : extras.keySet()) {
      if (key.startsWith("properties_")) {
        properties.put(key.replaceAll("properties_", ""), extras.get(key));
      }
    }

    Timber.d("analytics.screen(%s, %s, %s);", category, name, properties);
    Analytics.with(this).screen(category, name, properties);
  }

  void identify(Intent intent) {
    String userId = intent.getStringExtra("userId");

    Bundle extras = intent.getExtras();
    Traits traits = new Traits();
    for (String key : extras.keySet()) {
      if (key.startsWith("traits_")) {
        traits.put(key.replaceAll("traits_", ""), extras.get(key));
      }
    }

    Timber.d("analytics.identify(%s, %s);", userId, traits);
    Analytics.with(this).identify(userId, traits, null);
  }

  void alias(Intent intent) {
    String userId = intent.getStringExtra("userId");

    Timber.d("analytics.alias(%s);", userId);
    Analytics.with(this).alias(userId);
  }

  void group(Intent intent) {
    String groupId = intent.getStringExtra("groupId");

    Bundle extras = intent.getExtras();
    Traits traits = new Traits();
    for (String key : extras.keySet()) {
      if (key.startsWith("traits_")) {
        traits.put(key.replaceAll("traits_", ""), extras.get(key));
      }
    }

    Timber.d("analytics.group(%s, %s);", groupId, traits);
    Analytics.with(this).group(groupId, traits, null);
  }

  void flush() {
    Timber.d("analytics.flush();");
    Analytics.with(this).flush();
  }

  void reset() {
    Timber.d("analytics.reset();");
    Analytics.with(this).reset();
  }
}
