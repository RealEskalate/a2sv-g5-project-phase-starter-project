import 'dart:io';

import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:image_picker/image_picker.dart';

import '../../../auth/domain/entity/auth_entity.dart';
import '../../../auth/presentation/bloc/auth_bloc.dart';
import '../../../auth/presentation/bloc/auth_event.dart';
import '../../../auth/presentation/bloc/auth_state.dart';
import '../../../auth/presentation/pages/login_page.dart';
import '../../../chat/presentation/pages/chat_screen.dart';
import 'add_page.dart';
import 'search.dart';

class SideBar extends StatefulWidget {
  const SideBar({super.key});

  @override
  _SideBarState createState() => _SideBarState();
}

class _SideBarState extends State<SideBar> {
  File? _image;

  Future<void> _pickImage() async {
    final picker = ImagePicker();
    final pickedFile = await picker.pickImage(source: ImageSource.gallery);

    if (pickedFile != null) {
      setState(() {
        _image = File(pickedFile.path);
      });
    }
  }

  @override
  Widget build(BuildContext context) {
    context.read<AuthBloc>().add(GetUserProfileEvent());

    return Drawer(
      child: ListView(
        padding: EdgeInsets.zero,
        children: <Widget>[
          BlocBuilder<AuthBloc, AuthState>(
            builder: (context, state) {
              return UserAccountsDrawerHeader(
                decoration: BoxDecoration(
                  color: const Color.fromRGBO(63, 81, 181, 1),
                ),
                accountName:
                    Text(state is UserProfileLoaded ? state.user.name : ''),
                accountEmail:
                    Text(state is UserProfileLoaded ? state.user.email : ''),
                currentAccountPicture: GestureDetector(
                  onTap: _pickImage,
                  child: CircleAvatar(
                    child: ClipOval(
                      child: _image != null
                          ? Image.file(
                              _image!,
                              fit: BoxFit.cover,
                              width: 90,
                              height: 90,
                            )
                          : Image.asset(
                              'images/shoes.png',
                              fit: BoxFit.cover,
                              width: 90,
                              height: 90,
                            ),
                    ),
                  ),
                ),
              );
            },
          ),
          ListTile(
            leading: const Icon(Icons.home),
            title: const Text('Home'),
            onTap: () {
              Navigator.pop(context);
            },
          ),
          BlocBuilder<AuthBloc, AuthState>(
            builder: (context, state) {
              return ListTile(
                leading: const Icon(Icons.add_circle_rounded),
                title: const Text('Add Product'),
                onTap: () {
                  Navigator.push(context, MaterialPageRoute(builder: (context) {
                    return AddPage(
                        user: state is UserProfileLoaded
                            ? state.user
                            : UserEntity(
                                id: '',
                                name: 'user',
                                email: 'user@gmail.com',
                              ));
                  }));
                },
              );
            },
          ),
          ListTile(
            leading: const Icon(Icons.search),
            title: const Text('Search'),
            onTap: () {
              Navigator.push(context, MaterialPageRoute(builder: (context) {
                return SearchPage();
              }));
            },
          ),
          ListTile(
            leading: const Icon(Icons.chat),
            title: const Text('Chat'),
            onTap: () {
              Navigator.push(context, MaterialPageRoute(builder: (context) {
                return ChatScreen();
              }));
            },
          ),
          Divider(),
          BlocListener<AuthBloc, AuthState>(
            listener: (context, state) {
              if (state is LogoutSuccess) {
                ScaffoldMessenger.of(context).showSnackBar(const SnackBar(
                  backgroundColor: Colors.green,
                  content: Text('Logout successfully'),
                ));
                Navigator.of(context).push(
                  MaterialPageRoute(
                    builder: (context) => LoginPage(
                      text: '',
                    ),
                  ),
                );
              } else if (state is AuthError) {
                ScaffoldMessenger.of(context).showSnackBar(
                    const SnackBar(content: Text('Logout failed')));
              }
              return;
            },
            child: ListTile(
              leading: const Icon(Icons.logout),
              title: const Text('Logout'),
              onTap: () {
                context.read<AuthBloc>().add(const LogoutEvent());
              },
            ),
          ),
        ],
      ),
    );
  }
}
