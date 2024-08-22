// import 'dart:io';

// import 'package:flutter/material.dart';
// import 'package:flutter_bloc/flutter_bloc.dart';
// import '../../data/models/product_model.dart'; // Import your ProductModel file
// import '../bloc/add_page/add_page_bloc.dart';
// import '../bloc/home_page/home_page_bloc.dart';

// class AddButton extends StatelessWidget {
//   final TextEditingController nameController;
//   final TextEditingController priceController;
//   final TextEditingController descriptionController;
//   final File? selectedImage;
//   final int? index;

//   AddButton({
//     required this.nameController,
//     required this.priceController,
//     required this.descriptionController,
//     this.selectedImage,
//     this.index,
//   });

//   @override
//   Widget build(BuildContext context) {
//     return SizedBox(
//       height: 50,
//       width: double.infinity,
//       child: ElevatedButton(
//         onPressed: () {
//           final productModel = ProductModel(
//             id: index?.toString() ?? '',
//             name: nameController.text,
//             price: double.parse(priceController.text),
//             description: descriptionController.text,
//             imageUrl: selectedImage?.path ?? '',
//           );

//           context.read<AddPageBloc>().add(AddProductEvent(productModel, selectedImage?.path ?? ''));
//           context.read<HomePageBloc>().add(AddProductToHomePageEvent(productModel));

//           Navigator.pop(context, productModel);
//         },
//         style: ElevatedButton.styleFrom(
//           backgroundColor: const Color.fromARGB(255, 54, 104, 255),
//           foregroundColor: Colors.white,
//           shape: RoundedRectangleBorder(
//             borderRadius: BorderRadius.circular(12),
//           ),
//         ),
//         child: const Text('ADD'),
//       ),
//     );
//   }
// }
