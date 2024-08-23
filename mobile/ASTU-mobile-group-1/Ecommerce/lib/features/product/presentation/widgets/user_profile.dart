import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:intl/intl.dart';

import '../../../../config/route/route.dart' as route;
import '../../../../core/cubit/user_cubit.dart';
import '../../../auth/domain/entities/user_entity.dart';
import '../../../auth/presentation/bloc/auth_bloc.dart';
import '../../../auth/presentation/bloc/auth_event.dart';
import '../../../auth/presentation/widgets/custom_text.dart';
import 'custom_image_container.dart';

Widget userProfile(context) {
  return Row(
    children: [
      GestureDetector(
        onTap: () {
          showModalBottomSheet(
            context: context,
            shape: const RoundedRectangleBorder(
              borderRadius: BorderRadius.vertical(top: Radius.circular(20)),
            ),
            builder: (context) {
              return Padding(
                padding: const EdgeInsets.all(16.0),
                child: Column(
                  mainAxisSize: MainAxisSize.min,
                  children: [
                    const Text(
                      'Profile',
                      style: TextStyle(
                        fontSize: 18,
                        fontWeight: FontWeight.bold,
                      ),
                    ),
                    const SizedBox(height: 20),
                    Row(
                      children: [
                        const CircleAvatar(
                          radius: 30,
                          backgroundImage:
                              AssetImage('lib/assets/images/profile.jpg'),
                        ),
                        const SizedBox(width: 10),
                        Column(
                          crossAxisAlignment: CrossAxisAlignment.start,
                          children: [
                            BlocBuilder<UserCubit, UserEntity?>(
                              builder: (context, user) {
                                return Text(
                                  user?.name ?? 'Name',
                                  style: const TextStyle(
                                    fontSize: 16,
                                    fontWeight: FontWeight.bold,
                                  ),
                                );
                              },
                            ),
                            BlocBuilder<UserCubit, UserEntity?>(
                              builder: (context, user) {
                                return Text(
                                  user?.email ?? 'Email',
                                  style: const TextStyle(
                                    fontSize: 14,
                                    color: Colors.grey,
                                  ),
                                );
                              },
                            ),
                          ],
                        ),
                      ],
                    ),
                    const SizedBox(height: 20),
                    ListTile(
                      leading: const Icon(Icons.logout),
                      title: const Text('Logout'),
                      onTap: () {
                        context.read<AuthBloc>().add(LogOutEvent());
                        Navigator.pop(context);
                        Navigator.pushReplacementNamed(
                            context, route.signInPage);
                      },
                    ),
                  ],
                ),
              );
            },
          );
        },
        child: const CustomImageContainer(
          imagePath: 'lib/assets/images/profile.jpg',
          height: 50,
          width: 50,
          borderRadius: BorderRadius.all(
            Radius.circular(11),
          ),
        ),
      ),
      const SizedBox(
        width: 10,
      ),
      Column(
        children: [
          Align(
            alignment: Alignment.centerLeft,
            child: CustomText(
              text: DateFormat('MMMM dd, yyyy').format(DateTime.now()),
              fontSize: 12,
              color: Colors.grey,
              textAlign: TextAlign.left,
            ),
          ),
          Row(
            children: [
              const CustomText(
                text: 'Hello ',
                fontSize: 15,
                color: Colors.grey,
              ),
              BlocBuilder<UserCubit, UserEntity?>(
                builder: (context, state) {
                  return CustomText(
                    text: state == null ? '' : state.name,
                    fontSize: 15,
                    fontWeight: FontWeight.bold,
                  );
                },
              )
            ],
          )
        ],
      )
    ],
  );
}
