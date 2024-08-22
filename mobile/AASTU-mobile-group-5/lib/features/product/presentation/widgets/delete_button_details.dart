import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

import '../../domain/use_case/delete_product.dart';
import '../bloc/details_page/details_page_bloc.dart';

class DeleteButtonDetails extends StatelessWidget {
  final String id; // Use String type

  const DeleteButtonDetails({super.key, required this.id});

  @override
  Widget build(BuildContext context) {
    return OutlinedButton(
      onPressed: () {
        BlocProvider.of<DetailsPageBloc>(context).add(DeleteDetailsEvent(DeleteProductParams(id)));
      },
      style: OutlinedButton.styleFrom(
        side: const BorderSide(color: Colors.red),
        shape: RoundedRectangleBorder(
          borderRadius: BorderRadius.circular(10),
        ),
      ),
      child: const Text(
        'DELETE',
        style: TextStyle(color: Colors.red),
      ),
    );
  }
}
