import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

import 'features/authentication/presentation/bloc/blocs.dart';
import 'features/chat/presentation/pages/chat_list.dart';
import 'features/chat/presentation/pages/chat_page.dart';
import 'features/chat/presentation/widgets/chat.dart';
import 'features/chat/presentation/widgets/chat_card.dart';
import 'features/chat/presentation/widgets/sideBar.dart';
import 'features/chat/presentation/widgets/user_avater.dart';
import 'features/chat/presentation/widgets/voice_chat.dart';
import 'features/product/presentation/bloc/blocs.dart';
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
      ],
      child: MaterialApp(
        debugShowCheckedModeBanner: false,
        home: Scaffold(
          key: scaffoldKey, // Assign the key to the Scaffold
          appBar: AppBar(
            title: Text('My Chat App'),
            leading: IconButton(
              icon: Icon(Icons.menu),
              onPressed: () {
                scaffoldKey.currentState?.openDrawer(); // Use the key to open the drawer
              },
            ),
          ),
          drawer: const Sidebar(), // Add your Sidebar here as the drawer
          body: const Column(
            children: [
              // Placeholder for the chat widgets or other contents
              // Chat(message: "Sample message"),
              // SizedBox(height: 20),
              // RealTimeAudioRecorder(),
            ],
          ),
        ),
      ),
    );
  }
}
