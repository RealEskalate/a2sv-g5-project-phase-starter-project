import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:google_fonts/google_fonts.dart';
import 'features/product/domain/entities/product.dart';
import 'features/product/domain/use_case/add_product.dart';
import 'features/product/domain/use_case/delete_product.dart';
import 'features/product/domain/use_case/get_product.dart';
import 'features/product/domain/use_case/update_product.dart';
import 'features/product/presentation/bloc/add_page/add_page_bloc.dart';
import 'features/product/presentation/bloc/details_page/details_page_bloc.dart';
import 'features/product/presentation/bloc/home_page/home_page_bloc.dart';
import 'features/product/presentation/bloc/search_page/search_page_bloc.dart';
import 'features/product/presentation/bloc/update_page/update_page_bloc.dart';
import 'features/product/presentation/pages/product_add_page.dart';
import 'features/product/presentation/pages/product_details_page.dart';
import 'features/product/presentation/pages/product_home_page.dart';
import 'features/product/presentation/pages/product_search_page.dart';
import 'features/product/presentation/pages/product_update_page.dart';
import 'features/product/presentation/widgets/navigation_animation.dart';
import 'features/product/presentation/widgets/theme_cubit.dart';
import 'features/user/data/data_sources/user_local_data_source.dart';
import 'features/user/presentation/bloc/authentication/authentication_bloc.dart';
import 'features/user/presentation/pages/sign_in.dart';
import 'features/user/presentation/pages/sign_up.dart';
import 'features/user/presentation/pages/splash_screen.dart';
import 'service_locator.dart';

Future<void> main() async {
  WidgetsFlutterBinding.ensureInitialized();
  await setupLocator();
  runApp(const MyApp());
}

class MyApp extends StatelessWidget {
  const MyApp({super.key});
  @override
  Widget build(BuildContext context) {
    final textTheme = GoogleFonts.poppinsTextTheme();
    final darkTextTheme = textTheme.apply(
      bodyColor: Colors.white,
      displayColor: Colors.white,
    );
    final localDataSource = getIt<UserLocalDataSource>();
    return MultiBlocProvider(
      providers: [
        BlocProvider<HomePageBloc>(
          create: (context) {
            final bloc = getIt<HomePageBloc>();
            bloc.add(FetchAllProductsEvent());
            return bloc;
          },
        ),
        BlocProvider<AuthenticationBloc>(
          create: (context) => AuthenticationBloc(localDataSource),
        ),
        BlocProvider<ThemeCubit>(
          create: (context) => ThemeCubit(),
        ),
        // Add other providers here if needed
      ],
      child: BlocBuilder<ThemeCubit, ThemeMode>(
        builder: (context, themeMode) {
          return MaterialApp(
            debugShowCheckedModeBanner: false,
            theme: ThemeData(
              textTheme: textTheme, // Apply the common text theme to the light theme
            ),
            // darkTheme: ThemeData.dark().copyWith(
            //   textTheme: darkTextTheme, // Apply the dark text theme to the dark theme
            // ),
             darkTheme : ThemeData.dark().copyWith(
      scaffoldBackgroundColor: const Color(0xFF121212), // Softer dark background color
      textTheme: darkTextTheme,
    ),
            themeMode: themeMode,
            title: 'Product App',
            initialRoute: '/splash',
            routes: {
              '/splash': (context) => const SplashScreen(),
              '/signin': (context) => const SignInPage(),
              '/signup': (context) => const SignUpPage(),
            },
            onGenerateRoute: (settings) {
              switch (settings.name) {
                case '/details':
                  final product = settings.arguments as Product;
                  return MaterialPageRoute(
                    builder: (context) {
                      return BlocProvider(
                        create: (context) => DetailsPageBloc(
                          getIt<GetProduct>(),
                          getIt<DeleteProduct>(),
                        )..add(FetchProductByIdEvent(
                            GetProductParams(product.id))),
                        child: DetailsPage(id: product.id),
                      );
                    },
                    settings: settings,
                  );
                case '/add':
                  return SlidePageRoute(
                    page: BlocProvider(
                      create: (context) => AddPageBloc(
                        getIt<AddProduct>(),
                      ),
                      child: const AddPage(),
                    ),
                  );
                case '/update':
                  final product = settings.arguments as Product;
                  return MaterialPageRoute(
                    builder: (context) {
                      return BlocProvider(
                        create: (context) => UpdatePageBloc(
                          getIt<UpdateProduct>(),
                          getIt<DeleteProduct>(),
                        ),
                        child: UpdatePage(product: product),
                      );
                    },
                    settings: settings,
                  );
                case '/search':
                  return ScalePageRoute(
                    page: BlocProvider(
                      create: (context) => getIt<SearchPageBloc>()
                        ..add(FetchAllProductsSearchEvent()),
                      child: const SearchPage(),
                    ),
                  );
                case '/home':
                  return MaterialPageRoute(
                    builder: (context) => const HomePage(),
                    settings: settings,
                  );
              }
              return null;
            },
          );
        },
      ),
    );
  }
}
