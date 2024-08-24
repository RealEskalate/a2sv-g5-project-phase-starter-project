import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

import 'features/product/presentation/screens/add_product_page.dart';
import 'features/product/presentation/screens/homepage.dart';
import 'features/product/presentation/screens/product_detail_page.dart';
import 'features/product/presentation/screens/search_product.dart';


void main() async {
  WidgetsFlutterBinding.ensureInitialized();
  await di.init();
  Bloc.observer = SimpleBlocObserver();

  runApp(MyApp());
}

class MyApp extends StatelessWidget {
  MyApp({super.key});
  final BlocMultiProvider blocMultiProvider = BlocMultiProvider();

  @override
  Widget build(BuildContext context) {
    return MultiBlocProvider(
      providers: blocMultiProvider.blocMultiProvider(),
      child: FutureBuilder<String>(
        future: AppRouter._determineInitialRoute(),
        builder: (context, snapshot) {
          if (snapshot.connectionState == ConnectionState.waiting) {
            return const CircularProgressIndicator();
          }

          final initialRoute = snapshot.data ?? '/signin_page';

          return MaterialApp(
            title: 'Flutter Demo',
            debugShowCheckedModeBanner: false,
            theme: ThemeData(
              colorScheme: ColorScheme.fromSeed(
                seedColor: const Color.fromARGB(255, 35, 51, 203),
              ),
              useMaterial3: true,
              textTheme: const TextTheme(
                bodyLarge: TextStyle(fontFamily: 'Poppins'),
                bodyMedium: TextStyle(fontFamily: 'Poppins'),
                bodySmall: TextStyle(fontFamily: 'Poppins'),
                displayLarge: TextStyle(fontFamily: 'Poppins'),
                displayMedium: TextStyle(fontFamily: 'Poppins'),
                displaySmall: TextStyle(fontFamily: 'Poppins'),
                headlineLarge: TextStyle(fontFamily: 'Poppins'),
                headlineMedium: TextStyle(fontFamily: 'Poppins'),
                headlineSmall: TextStyle(fontFamily: 'Poppins'),
                titleLarge: TextStyle(fontFamily: 'Poppins'),
                titleMedium: TextStyle(fontFamily: 'Poppins'),
                titleSmall: TextStyle(fontFamily: 'Poppins'),
                labelLarge: TextStyle(fontFamily: 'Poppins'),
                labelMedium: TextStyle(fontFamily: 'Poppins'),
                labelSmall: TextStyle(fontFamily: 'Poppins'),
              ),
            ),
            initialRoute: initialRoute,
            routes: {
              '/splash_page': (context) => const SplashScreen(),
              '/signin_page': (context) => SigninPage(),
              '/signup_page': (context) => SignupPage(),
              '/product_detail_page': (context) => const ProductDetailPage(),
              '/add_product_page': (context) {
                final args = ModalRoute.of(context)!.settings.arguments as Map;
                final user = args['user'];
                return AddProductPage(user: user);
              },
              '/homepage': (context) => const HomePage(),
              '/search_page': (context) => const SearchPage(),
              '/chat_list_page': (context) => const ChatListPage(),
            },
          );
        },
      ),
    );
  }
}

class AppRouter {
  static Future<String> _determineInitialRoute() async {
    final token = await _getToken();

    if (token != null && token.isNotEmpty) {
      return '/homepage';
    } else {
      return '/signin_page';
    }
  }

  static Future<String?> _getToken() async {
    final prefs = await SharedPreferences.getInstance();
    final token = prefs.getString('tokenKey');
    return token;
  }
}
