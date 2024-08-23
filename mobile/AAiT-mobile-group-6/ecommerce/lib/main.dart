import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'bloc_observer.dart';
import 'features/auth/presentation/screens/signin_page.dart';
import 'features/auth/presentation/screens/signup_page.dart';
import 'features/auth/presentation/screens/splash_screen.dart';
import 'features/chat/presentation/screens/chat_list_page.dart';
import 'features/product/presentation/bloc/product_bloc.dart';
import 'features/product/presentation/screens/add_product_page.dart';
import 'features/product/presentation/screens/homepage.dart';
import 'features/product/presentation/screens/product_detail_page.dart';
import 'features/product/presentation/screens/search_product.dart';
import 'injection_container.dart' as di;

void main() async {
  WidgetsFlutterBinding.ensureInitialized();
  await di.init();
  Bloc.observer = SimpleBlocObserver();
  runApp(const MyApp());
}

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  @override
  Widget build(BuildContext context) {
    return BlocProvider(
      create: (_) => di.sl<ProductBloc>(),
      child: MaterialApp(
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
        home: ChatListPage(),
        routes: {
          '/splash_page': (context) => const SplashScreen(),
          '/signin_page': (context) => SigninPage(),
          '/signup_page': (context) => SignupPage(),
          '/product_detail_page': (context) => const ProductDetailPage(),
          '/add_product_page': (context) => const AddProductPage(),
          '/homepage': (context) => const HomePage(),
          '/search_page': (context) => const SearchPage(),
          '/chat_list_page': (context) => ChatListPage(),
        },
      ),
    );
  }
}
