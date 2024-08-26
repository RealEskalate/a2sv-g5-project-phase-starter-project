import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'config/route/route.dart' as route;
import 'core/cubit/user_cubit.dart';
import 'features/auth/presentation/bloc/auth_bloc.dart';
import 'features/chat/presentation/bloc/chat_bloc.dart';
import 'features/product/presentation/bloc/product_bloc.dart';
import 'injection_container.dart';

class App extends StatelessWidget {
  const App({super.key});

  @override
  Widget build(BuildContext context) {
    return MultiBlocProvider(
      providers: [
        BlocProvider(
          create: (context) => sl<ProductBloc>(),
        ),
        BlocProvider(
          create: (context) => sl<ChatBloc>(),
        ),
        BlocProvider(
          create: (context) => sl<AuthBloc>(),
        ),
        BlocProvider(create: (context) => sl<UserCubit>())
      ],
      child: MaterialApp(
        theme: ThemeData(
          fontFamily: 'Poppins',
        ),
        onGenerateRoute: route.controller,
        initialRoute: route.chatHomeScreen,
      ),
    );
  }
}
