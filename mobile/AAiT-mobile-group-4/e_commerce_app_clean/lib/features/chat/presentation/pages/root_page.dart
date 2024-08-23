import 'package:flutter/material.dart';

import '../widgets/notification_screen.dart';

class RootPage extends StatefulWidget {
  const RootPage({super.key});

  @override
  State<RootPage> createState() => _RootPageState();
}

class _RootPageState extends State<RootPage> {
  int Startindex = 0;
  List<Widget> _children = [
    const ChatPage(),
    const Center(child: Text('Contacts')),
    const Center(child: Text('Calls')),
    const Center(child: Text('Profile')),
  ];
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: _children[Startindex],
      bottomNavigationBar: NavigationBarTheme(
        data: NavigationBarThemeData(
          indicatorColor: Colors.grey[300],
          indicatorShape: RoundedRectangleBorder(
            borderRadius: BorderRadius.circular(30),
          ),
          height: 60,
          labelTextStyle: WidgetStateProperty.all<TextStyle>(
            const TextStyle(
              fontSize: 12,
              fontWeight: FontWeight.w500,
              color: Colors.black,
            ),
          ),
        ),
        child: NavigationBar(
          animationDuration: const Duration(seconds: 1),
          backgroundColor: Colors.white,
          selectedIndex: Startindex,
          onDestinationSelected: (index) {
            setState(() {
              Startindex = index;
            });
          },
          destinations: const [
            NavigationDestination(
              selectedIcon: Icon(Icons.message),
              icon: Icon(Icons.message_outlined),
              label: 'Message',
            ),
            NavigationDestination(
              selectedIcon: Icon(Icons.contacts),
              icon: Icon(Icons.contacts_outlined),
              label: 'Contacts',
            ),
            NavigationDestination(
              selectedIcon: Icon(Icons.call),
              icon: Icon(Icons.call_outlined),
              label: 'Calls',
            ),
            NavigationDestination(
              selectedIcon: Icon(Icons.person),
              icon: Icon(Icons.person_outlined),
              label: 'Profile',
            )
          ],
        ),
      ),
    );
  }
}
