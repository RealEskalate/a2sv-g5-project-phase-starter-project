import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:intl/intl.dart';

import '../../../auth/presentation/bloc/auth_bloc.dart';
import '../../../auth/presentation/page/login_page.dart';
import '../bloc/product_bloc.dart';
import '../bloc/product_events.dart';
import '../widgets/product_widgets.dart';

class ProductListPage extends StatelessWidget {
  static String routes = '/product_list_page';
  final DateTime now = DateTime.now();
  

  ProductListPage({super.key});


  @override
  Widget build(BuildContext context) {
    BlocProvider.of<ProductBloc>(context).add(LoadAllProductEvents());
    BlocProvider.of<AuthBloc>(context).add(GetMeEvent());
    return BlocListener<AuthBloc, AuthState>(
      listener: (context, state) {
        if (state is LogoutSuccess) {
          Navigator.of(context).pushReplacement(
            MaterialPageRoute(
              builder: (context) => LoginPage(),
            ),
          );
        } else if (state is LogoutFailedState) {
          ScaffoldMessenger.of(context)
              .showSnackBar(SnackBar(content: Text(state.message)));
        }
      },
      child: Scaffold(
        body: RefreshIndicator(
          onRefresh: () async {
            await Future.delayed(const Duration(seconds: 1));
            BlocProvider.of<ProductBloc>(context).add(LoadAllProductEvents());
          },
          child: SafeArea(
            child: Column(
              children: [
                BlocBuilder<AuthBloc, AuthState>(
                  builder: (context, state) {
                    String name = 'Sir';
                    String day = '${DateFormat('MMMM').format(now)} ${now.day}, ${now.year}';
                    String email = '...@gmail.com';
                    if (state is GetMeSuccessState) {
                      name = state.user.name;
                      email = state.user.email;
                    }
                    return UserInfo(
                      iconPres: () {
                        showDialog(
                            context: context,
                            builder: (context) {
                              return AlertDialog(
                                title: const Text('Logout'),
                                content: const Text(
                                    'Are you sure you want to logout.'),
                                actions: [
                                  FillCustomButton(
                                      press: () {
                                        BlocProvider.of<AuthBloc>(context)
                                            .add(LogOutEvent());
                                        Navigator.pop(context);
                                      },
                                      label: 'Logout')
                                ],
                              );
                            });
                      },
                      userName: name,
                      day: day,
                    );
                  },
                ),
                const SearchNavigator(),
                const ProductListDisplayer(),
              ],
            ),
          ),
        ),
      ),
    );
  }
}
