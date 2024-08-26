import 'package:flutter/material.dart';

import '../../../chat/presentation/pages/chat_home_screen.dart';
import '../../../product/presentation/pages/home_page.dart';

class GeneralPage extends StatefulWidget {
  const GeneralPage({super.key});

  @override
  State<GeneralPage> createState() => _GeneralPageState();
}

class _GeneralPageState extends State<GeneralPage> {
   int selectedPage = 0;
  final screens = [const HomePage(),const ChatHomeScreen()];
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      
      bottomNavigationBar: NavigationBar(
        selectedIndex: selectedPage,
        onDestinationSelected: (value) {
          setState(() {
            selectedPage = value;
          });
        },
        destinations: const [
                    NavigationDestination(icon: Icon(Icons.home), label: 'home'),
                    NavigationDestination(icon: Icon(Icons.message), label: 'message')
                    ],
      
        ),
        body: screens[selectedPage],
     
    );
  }
}

