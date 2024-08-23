import 'dart:io';

import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:image_picker/image_picker.dart';

import '../../domain/entities/product_entity.dart';
import '../bloc/product_bloc.dart';
import '../widgets/components/styles/custom_button.dart';
import '../widgets/components/styles/text_field_styles.dart';
import '../widgets/components/styles/text_style.dart';

// ignore: must_be_immutable
class UpdatePage extends StatefulWidget {
  ProductEntity selectedProduct;
  UpdatePage({
    super.key,
    required this.selectedProduct,
  });

  @override
  State<UpdatePage> createState() => _UpdatePageState();
}

class _UpdatePageState extends State<UpdatePage> {
  File? selectedImage;

  Future pickImageFromGallery() async {
    final returnedImage =
        await ImagePicker().pickImage(source: ImageSource.gallery);
    if (returnedImage != null) {
      setState(() {
        selectedImage = File(returnedImage.path);
      });
    }
  }

  final TextEditingController _nameController = TextEditingController();

  final TextEditingController _categoryController = TextEditingController();

  final TextEditingController _priceController = TextEditingController();

  final TextEditingController _descriptionController = TextEditingController();

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: BlocListener<ProductBloc, ProductState>(
        listener: (context, state) {
          if (state is ProductUpdatedState) {
            ScaffoldMessenger.of(context).showSnackBar(const SnackBar(
              content: Text('Successfully Updated Product'),
            ));
            Navigator.pushNamed(context, '/home_page');
          } else if (state is ProductErrorState) {
            ScaffoldMessenger.of(context).showSnackBar(const SnackBar(
              content: Text('error'),
            ));
          }
        },
        child: Container(
          margin: const EdgeInsets.all(32),
          child: SingleChildScrollView(
            child: Column(
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
                Row(
                  crossAxisAlignment: CrossAxisAlignment.center,
                  children: [
                    IconButton(
                      onPressed: () {
                        Navigator.pop(context);
                      },
                      icon: const Icon(
                        Icons.arrow_back_ios_rounded,
                        color: Color.fromRGBO(63, 81, 243, 1),
                      ),
                    ),
                    const SizedBox(width: 80),
                    const CustomTextStyle(
                        name: 'Update Product',
                        weight: FontWeight.w500,
                        size: 16),
                  ],
                ),
                const SizedBox(height: 23),
                GestureDetector(
                  onTap: () {
                    pickImageFromGallery();
                  },
                  child: Container(
                    width: 366,
                    height: 190,
                    decoration: BoxDecoration(
                      borderRadius: BorderRadius.circular(16),
                      color: const Color.fromRGBO(243, 243, 243, 1),
                    ),
                    child: selectedImage == null
                        ? const Center(
                            child: Column(
                              mainAxisAlignment: MainAxisAlignment.center,
                              children: [
                                Icon(
                                  Icons.image_outlined,
                                  size: 48,
                                ),
                                SizedBox(
                                  height: 16,
                                ),
                                CustomTextStyle(
                                    name: 'upload image',
                                    weight: FontWeight.w500,
                                    size: 14)
                              ],
                            ),
                          )
                        : Image.file(selectedImage!),
                  ),
                ),
                const SizedBox(height: 16),
                const CustomTextStyle(
                  name: 'name',
                  weight: FontWeight.w500,
                  size: 14,
                ),
                const SizedBox(height: 8),
                CustomTextField(
                  controller: _nameController,
                  hint: widget.selectedProduct.name,
                ),
                const SizedBox(height: 16),
                const CustomTextStyle(
                  name: 'category',
                  weight: FontWeight.w500,
                  size: 14,
                ),
                const SizedBox(height: 8),
                CustomTextField(
                  controller: _categoryController,
                  hint: 'men\'s shoes',
                ),
                const SizedBox(height: 16),
                const CustomTextStyle(
                  name: 'price',
                  weight: FontWeight.w500,
                  size: 14,
                ),
                const SizedBox(height: 8),
                Stack(
                  children: [
                    CustomTextField(
                      hint: widget.selectedProduct.price.toString(),
                      controller: _priceController,
                    ),
                    const Positioned(
                      left: 290,
                      top: 16,
                      child: Icon(Icons.attach_money),
                    )
                  ],
                ),
                const SizedBox(height: 16),
                const CustomTextStyle(
                  name: 'description',
                  weight: FontWeight.w500,
                  size: 14,
                ),
                const SizedBox(height: 8),
                CustomTextField(
                  lines: 5,
                  controller: _descriptionController,
                  hint: widget.selectedProduct.description,
                ),
                const SizedBox(height: 32),
                CustomButton(
                  pressed: () {
                    ProductEntity updateProduct = ProductEntity(
                      description: _descriptionController.text,
                      id: widget.selectedProduct.id,
                      imageUrl: selectedImage!.path,
                      name: _nameController.text,
                      price: double.parse(_priceController.text),
                    );
                    
                    context
                        .read<ProductBloc>()
                        .add(UpdateProductEvent(product:  updateProduct));
                  },
                  name: 'UPDATE',
                  width: double.infinity,
                  height: 50,
                  fgcolor: const Color.fromRGBO(255, 255, 255, 1),
                  bgcolor: const Color.fromRGBO(63, 81, 243, 1),
                ),
                const SizedBox(height: 16),
                const CustomButton(
                  pressed: null,
                  name: 'DELETE',
                  width: double.infinity,
                  height: 50,
                  fgcolor: Color.fromARGB(230, 255, 19, 19),
                  bgcolor: Color.fromRGBO(255, 255, 255, 1),
                ),
              ],
            ),
          ),
        ),
      ),
    );
  }
}
