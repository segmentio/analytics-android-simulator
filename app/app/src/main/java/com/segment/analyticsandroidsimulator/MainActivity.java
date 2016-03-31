package com.segment.analyticsandroidsimulator;

import android.app.Activity;
import android.content.Intent;
import android.os.Bundle;
import com.segment.analytics.Analytics;

import static com.segment.analyticsandroidsimulator.Utils.stringEqual;

public class MainActivity extends Activity {
  @Override protected void onCreate(Bundle savedInstanceState) {
    super.onCreate(savedInstanceState);
    setContentView(R.layout.activity_main);

    Intent intent = getIntent();
    String type = intent.getStringExtra("type");
    if (stringEqual("track", type)) {
      track();
    } else if (stringEqual("flush", type)) {
      flush();
    } else {
      throw new IllegalArgumentException("invalid event type: " + type);
    }
  }

  void track() {
    Intent intent = getIntent();
    String event = intent.getStringExtra("event");
    Analytics.with(this).track(event);
  }

  void flush() {
    Analytics.with(this).flush();
  }
}
