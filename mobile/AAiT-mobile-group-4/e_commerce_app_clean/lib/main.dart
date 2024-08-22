import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

import 'features/authentication/presentation/bloc/auth_bloc.dart';
import 'features/authentication/presentation/pages/cover_page.dart';
import 'features/authentication/presentation/pages/sign_in_page.dart';
import 'features/authentication/presentation/pages/sign_up_page.dart';
import 'features/product/domain/entities/product_entity.dart';
import 'features/product/presentation/bloc/product_bloc.dart';
import 'features/product/presentation/pages/details_page.dart';
import 'features/product/presentation/pages/home_page.dart';
import 'features/product/presentation/pages/product_add_page.dart';
import 'features/product/presentation/pages/product_search_page.dart';
import 'features/product/presentation/pages/update_page.dart';
import 'injection_container.dart' as di;
import 'injection_container.dart';

void main() async {
  WidgetsFlutterBinding.ensureInitialized();
  await di.init();
  runApp(
    MultiBlocProvider(
      providers: [
        BlocProvider<AuthBloc>(
          create: (context) => sl<AuthBloc>(),
        ),
        BlocProvider<ProductBloc>(
          create: (context) => sl<ProductBloc>(),
        ),

      ],
      child: MaterialApp(
        theme: ThemeData(
          primaryColor: const Color.fromRGBO(63, 81, 243, 1),
          secondaryHeaderColor: const Color.fromARGB(230, 255, 19, 19),
          useMaterial3: false,
        ),
        initialRoute: '/cover_page',
        onGenerateRoute: (settings) {
          if(settings.name == '/sign_in_page'){
            return createRoute(SignInPage());
          }
          else if(settings.name == '/sign_up_page'){
            return createRoute(const SignUpPage());
          }
          else if(settings.name == '/cover_page'){
            return createRoute(const CoverPage());
          }
          else if (settings.name == '/home_page') {
            return createRoute(const Home());
          } else if (settings.name == '/product_add_page') {
            return createRoute(const AddProudctPage());
          } else if (settings.name == '/product_search_page') {
            return createRoute(const ProductSearchPage());
          } else if (settings.name == '/details_page') {
            return createRoute(DetailsPage(selectedProduct: settings.arguments as ProductEntity));
          } else if (settings.name == '/update_page') {
            return createRoute(UpdatePage(selectedProduct: settings.arguments as ProductEntity));
          }
          return null;
        },
        title: 'Flutter App',
        debugShowCheckedModeBanner: false,
      ),
    ),
  );
}
PageRouteBuilder createRoute(Widget page) {
  return PageRouteBuilder(
    pageBuilder: (context, animation, secondaryAnimation) => page,
    transitionsBuilder: (context, animation, secondaryAnimation, child) {
      const begin = Offset(0.0, 1.0);
      const end = Offset.zero;
      const curve = Curves.ease;

      var tween = Tween(begin: begin, end: end).chain(CurveTween(curve: curve));

      return SlideTransition(
        position: animation.drive(tween),
        child: child,
      );
    },
  );
}