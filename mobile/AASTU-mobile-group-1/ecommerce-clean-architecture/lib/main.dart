import 'package:ecommerce/features/auth/presentation/bloc/authbloc/auth_bloc.dart';
import 'package:ecommerce/features/chat_feature/chat/presentation/pages/chat_history.dart';
import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:shared_preferences/shared_preferences.dart';

import 'features/auth/presentation/login.dart';
import 'features/auth/presentation/register.dart';
import 'features/auth/presentation/splash.dart';
import 'features/chat_feature/chat/domain/usecase/initialize_chat.dart';
import 'features/chat_feature/chat/presentation/pages/direct_messages.dart';
import 'features/product/data/data_sources/local_data_source.dart';
import 'features/product/data/model/product_model.dart';
import 'features/product/domain/usecases/getallproduct.dart';
import 'features/product/presentation/bloc/getallproductbloc/bloc/product_bloc.dart';
import 'features/product/presentation/pages/homepage.dart';
import 'service_locator.dart';
import 'temp_dart.dart';

// void main() async {
//   WidgetsFlutterBinding.ensureInitialized();
//   await setup();

//   runApp(const MyApp());
// }

void main() async {
  WidgetsFlutterBinding.ensureInitialized();
  await setup();

  runApp(const Sample());
}


class Sample extends StatelessWidget {
  const Sample({super.key});

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      home: TempDart(initializeChat:getIt<InitializeChat>() ,),
    );
  }
}

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  @override
Widget build(BuildContext context) {
    return MultiBlocProvider(
      providers: [
        BlocProvider<ProductBloc>(
          create: (context) => getIt<ProductBloc>(),
        ),
        BlocProvider<UserBloc>(
          create: (context) => getIt<UserBloc>(),
        ),
      ],
      child: MaterialApp(
        debugShowCheckedModeBanner: false,
        title: 'Flutter Demo',
        theme: ThemeData(
          colorScheme: ColorScheme.fromSeed(
            seedColor: Color.fromARGB(255, 255, 255, 255),
          ),
          useMaterial3: true,
        ),
        home: const splash(),
        routes: {
          // Add your routes here
        },
      ),
    );
  }
}
