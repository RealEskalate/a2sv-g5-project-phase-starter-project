import 'dart:io';
import 'package:ecommerce_app_ca_tdd/features/product/presentation/bloc/add/add_bloc.dart';
import 'package:ecommerce_app_ca_tdd/features/product/presentation/bloc/detail/detail_bloc.dart';
import 'package:ecommerce_app_ca_tdd/features/product/presentation/bloc/update/bloc/update_bloc.dart';
import 'package:ecommerce_app_ca_tdd/features/user_auth/presentation/bloc/get_user/get_user_bloc.dart';
import 'package:ecommerce_app_ca_tdd/features/user_auth/presentation/bloc/get_user/get_user_event.dart';
import 'package:ecommerce_app_ca_tdd/features/user_auth/presentation/bloc/login/bloc/sign_in_bloc.dart';
import 'package:ecommerce_app_ca_tdd/features/user_auth/presentation/bloc/signup/bloc/sign_up_bloc.dart';
import 'package:ecommerce_app_ca_tdd/features/user_auth/presentation/pages/login_screen.dart';
import 'package:ecommerce_app_ca_tdd/features/user_auth/presentation/pages/logout_screen.dart';
import 'package:ecommerce_app_ca_tdd/features/user_auth/presentation/pages/sign_up_screen.dart';
import 'package:ecommerce_app_ca_tdd/features/user_auth/presentation/pages/splash_screen.dart';
import 'package:flutter/material.dart';

import 'package:ecommerce_app_ca_tdd/features/product/presentation/pages/add_update.dart';
import 'package:ecommerce_app_ca_tdd/features/product/presentation/pages/details.dart';
import 'package:ecommerce_app_ca_tdd/features/product/presentation/pages/home.dart';
import 'package:ecommerce_app_ca_tdd/features/product/presentation/pages/search.dart';
import 'package:ecommerce_app_ca_tdd/features/product/presentation/pages/update.dart';
import 'package:get/get.dart';
import 'package:google_fonts/google_fonts.dart';
import 'package:page_transition/page_transition.dart';
import 'package:shared_preferences/shared_preferences.dart';
import 'features/product/data/data_sources/local_data_source/local_data_source.dart';
import 'features/product/data/data_sources/remote_data_source/remote_data_source.dart';
import 'features/product/data/models/product_models.dart';
// import 'models/product.dart';
import 'package:provider/provider.dart';

import 'features/product/domain/entities/product_entity.dart';
import 'features/product/domain/repositories/product_repository.dart';
import 'features/product/domain/usecases/get_all_usecase.dart';
import 'features/product/domain/usecases/get_detail_usecase.dart';
import 'features/product/presentation/bloc/home_bloc.dart';
import 'features/product/presentation/bloc/home_event.dart';
import 'models/product.dart';
import 'package:http/http.dart' as http;

import 'package:flutter_bloc/flutter_bloc.dart';
import 'locator.dart';

void main() async {
  WidgetsFlutterBinding.ensureInitialized();
  await init();
  Bloc.observer = CustomBlocObserver();
  runApp(
    const Main(),
  );
}

class Main extends StatelessWidget {
  const Main({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    
    return BlocProvider(
      create: (context) => sl.get<HomeBloc>()..add(GetProductsEvent()),
      child: MaterialApp(
        debugShowCheckedModeBanner: false,
        // initial for splash page
        initialRoute: '/splash',

        onGenerateRoute: (settings) {
          if (settings.name == '/detail') {
            final item = settings.arguments as ProductModel;
            return MaterialPageRoute(
              builder: (context) {
                return BlocProvider(
                  create: (context) => sl.get<DetailBloc>(),
                  child: DetailsPage(item: item),
                );
              },
            );
          } else if (settings.name == '/update') {
            final item = settings.arguments as ProductModel;
            return MaterialPageRoute(
              builder: (context) {
                return BlocProvider(
                  create: (context) => sl.get<UpdateBloc>(),
                  child: UpdatePage(product: item),
                );
              },
            );
          }
          // Handle other routes here
          return null;
        },
        routes: {
          '/home': (context) => MultiBlocProvider(
            providers: [
        BlocProvider(create: (context)=> sl.get<HomeBloc>()..add(GetProductsEvent())),
        BlocProvider(create: (context)=>sl.get<GetUserBloc>()..add(GetUserInfoEvent()))
            ],
            child: HomePage(),
          ),
          '/login': (context) => BlocProvider(
                create: (context) => sl.get<LoginBloc>(),
                child: LoginScreen(),
              ),
          '/signup': (context) => BlocProvider(
                create: (context) => sl.get<SignUpBloc>(),
                child: SignUpScreen(),
              ),
            '/logout': (context)=> LogoutScreen(),

          '/add': (context) => BlocProvider(
                create: (context) => sl.get<addBloc>(),
                child: AddUpdate(),
              ),
          '/search': (context) => BlocProvider(
              create: (context) => sl.get<HomeBloc>()..add(GetProductsEvent())),

          '/splash': (context) => SplashScreen(),
          // '/detail': (context) => DetailsPage(),
        },
      ),
    );
  }
}
