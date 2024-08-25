import 'package:flutter/material.dart';

class Sidebar extends StatelessWidget {
  const Sidebar({super.key});

  @override
  Widget build(BuildContext context) {
    return Drawer(
      child: ListView(
        padding: EdgeInsets.zero,
        children: <Widget>[
          // Drawer Header with Profile Information
          UserAccountsDrawerHeader(
            accountName: Text('Sabila Sayma'),
            accountEmail: Text('sabila@example.com'),
            currentAccountPicture: CircleAvatar(
              backgroundImage: AssetImage('assets/images/avater.png'),
            ),
            decoration: BoxDecoration(
              color: Colors.blue,
            ),
          ),

          // Home Page Navigation
          ListTile(
            leading: Icon(Icons.home),
            title: Text('Home'),
            onTap: () {
              Navigator.of(context).pushNamed('/home');
            },
          ),

          // Chats Navigation
          ListTile(
            leading: Icon(Icons.chat),
            title: Text('Chats'),
            onTap: () {
              Navigator.of(context).pushNamed('/chats');
            },
          ),

          // Dark Theme Toggle
          ListTile(
            leading: Icon(Icons.brightness_6),
            title: Text('Dark Theme'),
            trailing: Switch(
              value: Theme.of(context).brightness == Brightness.dark,
              onChanged: (value) {
                // Add your dark theme toggle logic here
              },
            ),
          ),

          // Logout
          ListTile(
            leading: Icon(Icons.logout),
            title: Text('Logout'),
            onTap: () {
              // Add your logout logic here
            },
          ),
        ],
      ),
    );
  }
}
