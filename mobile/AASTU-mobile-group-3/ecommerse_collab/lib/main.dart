import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

import 'features/authentication/presentation/bloc/blocs.dart';
import 'features/authentication/presentation/pages/onboarding.dart';
import 'features/chat/presentation/bloc/blocs.dart';
import 'features/chat/presentation/pages/chat_list.dart';
import 'features/chat/presentation/pages/chat_page.dart';
import 'features/chat/presentation/widgets/chat.dart';
import 'features/chat/presentation/widgets/chat_card.dart';
import 'features/chat/presentation/widgets/sideBar.dart';
import 'features/chat/presentation/widgets/user_avater.dart';
import 'features/chat/presentation/widgets/voice_chat.dart';
import 'features/product/presentation/bloc/blocs.dart';
import 'features/product/presentation/pages/home_page.dart';
import 'service_locator.dart';

// Create a GlobalKey for the Scaffold
final GlobalKey<ScaffoldState> scaffoldKey = GlobalKey<ScaffoldState>();

void main() async {
  WidgetsFlutterBinding.ensureInitialized();

  await setUp();
  runApp(const MyApp());
}

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  @override
  Widget build(BuildContext context) {
    var userBloc = getIt<UserBloc>();
    return MultiBlocProvider(
      providers: [
        BlocProvider(create: (context) => getIt<ProductBloc>()),
        BlocProvider(create: (context) => userBloc),
        BlocProvider(create: (context) => ChatBloc()),

      ],
      child: MaterialApp(
         
        debugShowCheckedModeBanner: false,
        home: Onboarding()
      ),
    );
  }
}
