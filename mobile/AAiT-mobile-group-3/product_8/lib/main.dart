import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:google_fonts/google_fonts.dart';

import 'bloc_observer.dart';
import 'features/auth/presentation/bloc/auth_bloc/auth_bloc.dart';
import 'features/auth/presentation/pages/sign_in_page.dart';
import 'features/auth/presentation/pages/sign_up_page.dart';
import 'features/auth/presentation/pages/splash_page.dart';
import 'features/product/domain/entities/product_entity.dart';
import 'features/product/presentation/bloc/product_bloc.dart';
import 'features/product/presentation/pages/add_product_page.dart';
import 'features/product/presentation/pages/home_page.dart';
import 'features/product/presentation/pages/product_details_page.dart';
import 'features/product/presentation/pages/search_page.dart';
import 'features/product/presentation/pages/update_page.dart';
import 'injection_container.dart' as di;
import 'injection_container.dart';

void main() async {
  WidgetsFlutterBinding.ensureInitialized();
  await di.init();
  Bloc.observer = SimpleBlocObserver();
  runApp(const MyApp());
}

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  // This widget is the root of your application.
  @override
  Widget build(BuildContext context) {
    return MultiBlocProvider(
      providers: [
        BlocProvider(
          create: (context) => ProductBloc(
              getProductByIdUsecase: sl(),
              getProductsUsecase: sl(),
              deleteProductUsecase: sl(),
              updateProductUsecase: sl(),
              insertProductUsecase: sl()),
        ),
        BlocProvider(
          create: (context) => AuthBloc(
           sl(),
            sl(),
            sl(),
            sl(),
          ),
        ),

         
      ],
      child: MaterialApp(
        onGenerateRoute: (settings) {
          switch (settings.name) {
            case '/splash':
              return _createRoute(const SplashPage());
            case '/signin':
              return _createRoute(const SignInPage());
            case '/signup':
              return _createRoute(const SignUpPage());
            case '/home':
              return _createRoute(const HomePage());
            case '/search':
              return _createRoute(const SearchPage());
            case '/detail/update':
              final args = settings.arguments as Product;
              return _createRoute(UpDate(
                product: args,
              ));
            case '/add':
              return _createRoute(const ADDPage());

            case '/detail':
              final args = settings.arguments as Product;
              return _createRoute(Detailspage(
                product: args,
              ));
            default:
              return null;
          }
        },
        initialRoute: '/splash',
        title: 'Flutter Demo',
        theme: ThemeData(
          textTheme: GoogleFonts.poppinsTextTheme(),
          primaryColor: const Color.fromARGB(255, 63, 81, 243),
          colorScheme: ColorScheme.fromSeed(
              seedColor: const Color.fromARGB(255, 63, 81, 243)),
          useMaterial3: true,
        ),
        // home: const HomePage(),
        debugShowCheckedModeBanner: false,
      ),
    );
  }
}

PageRouteBuilder _createRoute(Widget page) {
  return PageRouteBuilder(
    pageBuilder: (context, animation, secondaryAnimation) => page,
    transitionsBuilder: (context, animation, secondaryAnimation, child) {
      const begin = Offset(1.0, 0.0);
      const end = Offset.zero;
      const curve = Curves.easeIn;

      var tween = Tween(begin: begin, end: end).chain(CurveTween(curve: curve));

      return SlideTransition(
        position: animation.drive(tween),
        child: child,
      );
    },
  );
}
