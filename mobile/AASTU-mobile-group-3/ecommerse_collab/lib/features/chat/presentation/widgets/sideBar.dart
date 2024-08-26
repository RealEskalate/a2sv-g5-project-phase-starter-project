import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

import '../../../authentication/domain/entity/user.dart';
import '../../../authentication/presentation/bloc/blocs.dart';
import '../../../authentication/presentation/bloc/events.dart';
import '../../../product/presentation/pages/home_page.dart';
import '../pages/chat_list.dart';
import 'chat.dart';

class Sidebar extends StatelessWidget {
  final User user;

  const Sidebar({super.key, required this.user });


  @override
  Widget build(BuildContext context) {
    return Drawer(
      child: ListView(
        padding: EdgeInsets.zero,
        children: <Widget>[
          // Drawer Header with Profile Information
          UserAccountsDrawerHeader(
            accountName: Text(user.username),
            accountEmail: Text(user.email),
            currentAccountPicture: CircleAvatar(
              backgroundImage: AssetImage('assets/images/avater.png'),
            ),
            decoration: BoxDecoration(
              color:Color(0xFF3F51F3),
            ),
          ),

          // Home Page Navigation
          ListTile(
            leading: Icon(Icons.home),
            title: Text('Home'),
            onTap: () {
              Navigator.of(context)
                    .push(MaterialPageRoute(builder: (BuildContext context){
                      return HomePage(user: user);
                    }),
              );
            },
          ),

          // Chats Navigation
          ListTile(
            leading: Icon(Icons.chat),
            title: Text('Chats'),
            onTap: () {
              Navigator.of(context)
                    .push(MaterialPageRoute(builder: (BuildContext context){
                      return ChatList(user);
                    }),
                    );
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
             
              context.read<UserBloc>().add(LogOutEvent());
            },
          ),
        ],
      ),
    );
  }
}
