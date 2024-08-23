import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

import '../../../../../../core/Text_Style/text_style.dart';
import '../../../state/user_states/login_user_states_bloc.dart';
import '../../../state/user_states/login_user_states_event.dart';
import '../../../state/user_states/login_user_states_state.dart';
import '../../seachProduct/apply_filter.dart';

class ProfilePage extends StatelessWidget {
  const ProfilePage({super.key});

  @override
  Widget build(BuildContext context) {
    return Container(
      color: Colors.grey,
      padding: const EdgeInsets.all(10),
      height: 200,
      child: Column(
        mainAxisAlignment: MainAxisAlignment.spaceBetween,
        children: [
          TextStyles(
            text: 'Profile',
            fontColor: Colors.black,
            fontSizes: 20,
            fontWeight: FontWeight.bold,
            fontFamily: 'CaveatBrush',
          ),
          Row(
            children: [
              TextStyles(
                text: 'Name: ',
                fontColor: Colors.black,
                fontSizes: 25,
                fontWeight: FontWeight.bold,
                fontFamily: 'Popins',
              ),
              BlocBuilder<LoginUserStatesBloc, LoginUserStates>(
                builder: (context, state) {
                  final name = state is ProfileDetailState ? state.name : '';

                  return TextStyles(
                    text: name,
                    fontColor: Colors.black,
                    fontSizes: 20,
                    fontWeight: FontWeight.normal,
                    fontFamily: 'CaveatBrush',
                  );
                },
              ),
            ],
          ),
          Row(
            children: [
              TextStyles(
                text: 'Email: ',
                fontColor: Colors.black,
                fontSizes: 25,
                fontWeight: FontWeight.normal,
                fontFamily: 'Popins',
              ),
              BlocBuilder<LoginUserStatesBloc, LoginUserStates>(
                builder: (context, state) {
                  final email = state is ProfileDetailState ? state.email : '';

                  return TextStyles(
                    text: email,
                    fontColor: Colors.black,
                    fontSizes: 20,
                    fontWeight: FontWeight.bold,
                    fontFamily: 'CaveatBrush',
                  );
                },
              ),
            ],
          ),
         GestureDetector(
            
            onTap: () {
             
              context
                  .read<LoginUserStatesBloc>()
                  .add(LogedOutUserStatesEvent());
              
              
            Navigator.pushReplacementNamed(context, '/login');
              
              
            },
            child: const ApplyFilter(
              text: 'LogOut',
            ))
        ],
      ),
    );
  }
}
