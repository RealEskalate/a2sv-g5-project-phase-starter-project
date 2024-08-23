import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

import 'core/themes/themes.dart';
import 'dependency_injection.dart';
import 'features/auth/presentation/bloc/auth_bloc.dart';
import 'features/auth/presentation/bloc/cubit/user_input_validation_cubit.dart';
import 'features/auth/presentation/page/login_page.dart';
import 'features/auth/presentation/page/signup_page.dart';
import 'features/product/presentation/bloc/cubit/input_validation_cubit.dart';
import 'features/product/presentation/bloc/product_bloc.dart';
import 'features/product/presentation/pages/add_product_page.dart';
import 'features/product/presentation/pages/product_list_page.dart';
import 'features/product/presentation/pages/search_product_page.dart';
import 'features/product/presentation/pages/single_product_page.dart';
import 'features/product/presentation/pages/update_product_page.dart';
import 'splash_page.dart';

void main() async {
  WidgetsFlutterBinding().ensureSemantics();
  await init();
  runApp(const MyApp());
}

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  @override
  Widget build(BuildContext context) {
    return MultiBlocProvider(
      providers: [
        BlocProvider(
          create: (_) => locator<ProductBloc>(),
        ),
        BlocProvider(
          create: (_) => locator<InputValidationCubit>(),
        ),
        BlocProvider(
          create: (_) => locator<AuthBloc>(),
        ),
        BlocProvider(
          create: (_) => locator<UserInputValidationCubit>(),
        )
      ],
      child: MaterialApp(
        theme: MyTheme.lightTheme,
        routes: {
          AddProductPage.routes: (context) => AddProductPage(),
          ProductListPage.routes: (context) => const ProductListPage(),
          SingleProduct.routes: (context) => SingleProduct(),
          SearchProduct.routes: (context) => const SearchProduct(),
          SplashPage.routes: (context) => const SplashPage(),
          LoginPage.routes: (context) => LoginPage(),
          SignUpPage.routes: (context) => SignUpPage(),
          UpdateProductPage.routes: (context) => UpdateProductPage()
        },
        initialRoute: SplashPage.routes,
      ),
    );
  }
}
