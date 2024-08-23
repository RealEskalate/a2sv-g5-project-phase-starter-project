import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:flutter_easyloading/flutter_easyloading.dart';

import '../../state/input_button_activation/bottum_state.dart';
import '../../state/input_button_activation/button_bloc.dart';
import '../../state/input_button_activation/button_event.dart';
import 'add_delete_button.dart';


class ConvertInputToJson extends StatelessWidget {
  final Map<String, dynamic> data;
  const ConvertInputToJson({super.key, required this.data});

  @override
  Widget build(BuildContext context) {
    
    return BlocBuilder<ButtonBloc, BottumState>(
      builder: (context, state) {
         final bool check = state is OnButtonActivate ? state.isActivate : false;
         final bool type = data['id'].isEmpty;
        return GestureDetector(
          key: const Key('add_button'),
          onTap: check?() {
            EasyLoading.showProgress(0.3, status: type?'Uploading...':'Updating...');
            context.read<ButtonBloc>().add(type?AddProductEvent():UpdateProductEvent());
          }:null,
          child: AddDeleteButton(
            color: check?Colors.blue:Colors.grey,
            text: data['name'] != '' ? 'EDIT' : 'ADD',
            borderColor:check?Colors.blue:Colors.grey,
          ),
        );
      },
    );
  }
}
