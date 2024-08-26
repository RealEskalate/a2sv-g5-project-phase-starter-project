import 'package:flutter/material.dart';

import 'app.dart';
import 'injection_container.dart' as sl;

Future<void> main() async {
  WidgetsFlutterBinding.ensureInitialized();
  await sl.init();

  runApp(const App());
}
