import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

import '../../../../authentication/presentation/bloc/auth_bloc.dart';
import 'styles/text_style.dart';

class HeaderView extends StatelessWidget {
  const HeaderView({super.key});

  @override
  Widget build(BuildContext context) {
    return Column(children: [
      SizedBox(
        width: MediaQuery.of(context).size.width,
        height: 50,
        child: Row(
          mainAxisAlignment: MainAxisAlignment.spaceBetween,
          children: [
            SizedBox(
              child: Row(
                crossAxisAlignment: CrossAxisAlignment.center,
                children: [
                  Container(
                    width: 50,
                    height: 50,
                    decoration: BoxDecoration(
                      borderRadius: BorderRadius.circular(11),
                      color: Theme.of(context).primaryColor,
                    ),
                  ),
                  const SizedBox(width: 10),
                  Column(
                    mainAxisAlignment: MainAxisAlignment.center,
                    crossAxisAlignment: CrossAxisAlignment.start,
                    children: [
                      const CustomTextStyle(
                        name: 'July 14,2023',
                        weight: FontWeight.w400,
                        size: 12,
                        family: 'Syne',
                        color: Color.fromRGBO(170, 170, 170, 1),
                      ),
                      BlocBuilder<AuthBloc, AuthState>(
                        builder: (context, state) {
                          if (state is AuthUserLoaded) {
                            // final String name = state.userEntity.name;
                            return Row(
                              children: [
                                const CustomTextStyle(
                                    name: 'Hello, ',
                                    weight: FontWeight.w400,
                                    size: 15),
                                CustomTextStyle(
                                    name: state.userEntity.name,
                                    weight: FontWeight.w600,
                                    size: 15),
                              ],
                            );
                          } else {
                            return Container();
                          }
                        },
                      )
                    ],
                  )
                ],
              ),
            ),
            Row(children: [
              Container(
                width: 40,
                height: 40,
                decoration: BoxDecoration(
                  border: Border.all(
                    color: const Color.fromRGBO(221, 221, 221, 1),
                    width: 1.0,
                  ),
                  borderRadius: BorderRadius.circular(9.0),
                ),
                child: Stack(
                  children: [
                    IconButton(
                      icon: const Icon(Icons.notifications_none_outlined),
                      color: const Color.fromRGBO(102, 102, 102, 1),
                      onPressed: () {},
                    ),
                    Positioned(
                        left: 20,
                        top: 12,
                        child: Icon(
                          Icons.circle,
                          size: 8,
                          color: Colors.blue[800],
                        ))
                  ],
                ),
              ),
              const SizedBox(width: 10,),
              SizedBox(
                width: 42,
                height: 42,
                child: Material(
                  color: Theme.of(context).primaryColor,
                  shape: RoundedRectangleBorder(
                    borderRadius: BorderRadius.circular(10)
                  ),
                  child: IconButton(
                    iconSize: 20,
                    color: Colors.white,
                    icon: const Icon(Icons.logout),
                    onPressed: () {
                      // Handle logout logic here
                      showDialog(
                        context: context,
                        builder: (context) => AlertDialog(
                          title: const Text('Logout'),
                          content: const Text('Are you sure you want to logout?'),
                          shape: RoundedRectangleBorder(
                            borderRadius: BorderRadius.circular(24),
                          ),
                          actions: [
                            TextButton(
                              onPressed: () {
                                Navigator.pop(context);
                              },
                              child: const Text('Cancel'),
                            ),
                            TextButton(
                              onPressed: () {
                                context.read<AuthBloc>().add(LogOutEvent());
                                Navigator.popAndPushNamed(context, '/sign_in_page');
                              },
                              child: const Text('Logout'),
                            ),
                          ],
                        ),
                      );
                    },
                  ),
                ),
              ),
            ]),
          ],
        ),
      ),
      const SizedBox(height: 38.0),
      SizedBox(
        width: MediaQuery.of(context).size.width,
        child: Row(
          mainAxisAlignment: MainAxisAlignment.spaceBetween,
          children: [
            const CustomTextStyle(
                name: 'Available Products', weight: FontWeight.w600, size: 24),
            Container(
              width: 40,
              height: 40,
              decoration: BoxDecoration(
                border: Border.all(
                  color: const Color.fromRGBO(221, 221, 221, 1), // Border color
                  width: 1.0, // Border width
                ),
                borderRadius: BorderRadius.circular(9.0), // Border radius
              ),
              child: IconButton(
                  icon: const Icon(
                    Icons.search,
                    size: 24,
                  ),
                  color: const Color.fromRGBO(221, 221, 221, 1), // Icon color
                  onPressed: () {
                    Navigator.pushNamed(
                      context,
                      '/product_search_page',
                    );
                  }),
            ),
          ],
        ),
      ),
    ]);
  }
}
