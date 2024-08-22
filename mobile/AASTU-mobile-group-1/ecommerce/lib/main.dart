import 'package:ecommerce/core/import/import_file.dart';
import 'package:flutter/material.dart';

Future<void> main() async {
  WidgetsFlutterBinding.ensureInitialized();
  Bloc.observer = ProductObserver();

  await setUp();

  runApp(const MainPage());
}

class MainPage extends StatelessWidget {
  const MainPage({super.key});

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      debugShowCheckedModeBanner: false,
      initialRoute: '/',
      routes: {
        '/': (context) =>  Container(),
        '/add_product': (context) =>  Container(),
      },
      onGenerateRoute: (settings) {
        if (settings.name == '/update') {
          final item = settings.arguments as ProductEntity?;
          return MaterialPageRoute(
            builder: (context) {
              return Container();
            },
          );
        }
        return null;
      },
    );
  }
}
