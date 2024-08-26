// import 'package:ecommers/Features/AddProduct/Presentation/UI/addProduct.dart';

import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:flutter_easyloading/flutter_easyloading.dart';
import 'package:provider/provider.dart';

import 'core/utility/socket_impl.dart';
import 'features/chat/presentation/UI/Chat_page/IndividualChatPage.dart';
import 'features/chat/presentation/UI/chatScreen/chatapp.dart';
import 'features/chat/presentation/bloc/chat_bloc.dart';
import 'features/ecommerce/presentation/UI/add_product/add_product.dart';
import 'features/ecommerce/presentation/UI/home/Product_detail/detail_page.dart';
import 'features/ecommerce/presentation/UI/home/home.dart';
import 'features/ecommerce/presentation/UI/seachProduct/search_screen.dart';
import 'features/ecommerce/presentation/state/image_input_display/image_bloc.dart';
import 'features/ecommerce/presentation/state/input_button_activation/button_bloc.dart';

import 'features/ecommerce/presentation/state/product_bloc/product_bloc.dart';
import 'features/ecommerce/presentation/state/user_states/login_user_states_bloc.dart';
import 'features/login/presentation/UI/SplashScreen/splash_screen.dart';
import 'features/login/presentation/UI/login_home_page.dart';
import 'features/login/presentation/UI/register_body.dart';
import 'features/login/presentation/state/Login_Registration/login_registration_bloc.dart';
import 'features/login/presentation/state/login/login_bloc.dart';
import 'injection.dart' as di;

void main() async {
  WidgetsFlutterBinding.ensureInitialized();
  await di.setUpLocator();
  runApp(
   
      const MyApp(),
    
  );
}

class MyApp extends StatelessWidget {
  const MyApp({super.key});
  @override
  Widget build(BuildContext context) {
    return MultiProvider(
      providers: [
        
        BlocProvider<ProductBloc>(
          create: (context) => di.locator<ProductBloc>(),
        ),
        BlocProvider<ImageBloc>(
          create: (context) => di.locator<ImageBloc>(),
        ),
        BlocProvider<ButtonBloc>(
          create: (context) => di.locator<ButtonBloc>(),
        ),
        BlocProvider<LoginBloc>(
          create: (context) => di.locator<LoginBloc>(),
        ),
        BlocProvider<LoginRegistrationBloc>(
          create: (context) => di.locator<LoginRegistrationBloc>(),
        ),
        BlocProvider<LoginUserStatesBloc>(
          create: (context) => di.locator<LoginUserStatesBloc>(),),
        BlocProvider<ChatBloc>(
          create: (context) => di.locator<ChatBloc>(),)
        // Add more BlocProviders here if needed
      ],
      child: MaterialApp(
        title: 'Ecommers',
        initialRoute: '/',
        routes: {
          '/add-product': (context) => const AddProduct(),
          '/detail': (context) => const DetailPage(),
          '/search': (context) => const SearchScreen(),
          '/home': (context) => const HomeScreen(),
          '/login': (context) => const LoginHomePage(),
          '/registration': (context) => const RegisterBody(),
          '/chat-message': (context) => const IndividualChatPage(),
          '/chat-screen': (context) =>  ChatListPage(),
        },
        debugShowCheckedModeBanner: false,
        theme: ThemeData(
          colorScheme: ColorScheme.fromSeed(
              seedColor: const Color.fromARGB(255, 255, 255, 255)),
          useMaterial3: true,
        ),
        home: const SplashScreen(),
        builder: EasyLoading.init(),
      ),
    );
  }
}
