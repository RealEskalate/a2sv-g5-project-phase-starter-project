import 'dart:io';

import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:image_picker/image_picker.dart';

import '../../domain/entities/product_entity.dart';
import '../bloc/product_bloc.dart';
import '../bloc/product_event.dart';
import '../bloc/product_state.dart';
import '../widgets/custom_buttom.dart';

class UpDate extends StatefulWidget {
  final Product product;
  const UpDate({super.key , required this.product});
  static const routeName = '/detail/update';

  @override
  State<UpDate> createState() => _UpDateState();
  
}

class _UpDateState extends State<UpDate> {
  late TextEditingController nameController = TextEditingController();
  late TextEditingController typeController = TextEditingController();
  late TextEditingController descriptionController = TextEditingController();
  late TextEditingController priceController = TextEditingController();
  File? _image;

  final ImagePicker _picker = ImagePicker();
  @override
  void initState() {
    super.initState();

   nameController = TextEditingController(text: widget.product?.name ?? '');
      priceController =
          TextEditingController(text: widget.product?.price.toString() ?? '');
      descriptionController =
          TextEditingController(text: widget.product?.description ?? '');
      typeController = TextEditingController();
  }

  Future<void> _pickImage() async {
    final XFile? pickedFile =
        await _picker.pickImage(source: ImageSource.gallery);
    if (pickedFile != null) {
      setState(() {
        _image = File(pickedFile.path); // Store the selected image in _image
      });
    }
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: const Color.fromRGBO(254, 254, 254, 1),
      appBar: AppBar(
        leading: IconButton(
            onPressed: () => {Navigator.pop(context)},
            icon: Icon(
              Icons.arrow_back_ios_outlined,
              color: Colors.indigoAccent.shade400,
            )),
        title: const Text(
          'Update Product',
          style: TextStyle(fontWeight: FontWeight.bold),
        ),
        centerTitle: true,
      ),
      body: BlocListener<ProductBloc, ProductState>(
        listener: (context, state) {
          if (state is ProductUpdatedState) {
            ScaffoldMessenger.of(context).showSnackBar(
                const SnackBar(content: Text('Product Updated Successfully')));
            Navigator.of(context).pushNamed('/home');
          }  else if (state is ProductUpdatedErrorState) {
            ScaffoldMessenger.of(context)
                .showSnackBar(SnackBar(content: Text(state.message)));
          }
        },
        child: SingleChildScrollView(
          child: Padding(
            padding: const EdgeInsets.all(20.0),
            child: Column(
              mainAxisAlignment: MainAxisAlignment.start,
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
                GestureDetector(
                  onTap: _pickImage,
                  child: Container(
                      height: 200,
                      width: double.infinity,
                      decoration: BoxDecoration(
                        color: const Color.fromRGBO(243, 243, 243, 1),
                        borderRadius: BorderRadius.circular(8),
                        image:
                            _image != null // Show the picked image if available
                                ? DecorationImage(
                                    image: FileImage(_image!),
                                    fit: BoxFit.cover,
                                  )
                                : null,
                      ),
                       child: _image == null
                          ? const Column(
                              mainAxisAlignment: MainAxisAlignment.center,
                              children: [
                                Icon(
                                  Icons.image_outlined,
                                  size: 40,
                                ),
                                SizedBox(
                                  height: 10,
                                ),
                                Text('Upload Image')
                              ],
                            )
                          : null,
                    ),
                      ),
                
                const SizedBox(
                  height: 20.0,
                ),
                const Text(
                  'name',
                  textAlign: TextAlign.start,
                ),
                const SizedBox(
                  height: 10.0,
                ),
                TextField(
                  controller: nameController,
                  decoration: InputDecoration(
                    border: OutlineInputBorder(
                        borderSide: BorderSide.none,
                        borderRadius: BorderRadius.circular(10)),
                    filled: true,
                    fillColor: const Color.fromARGB(169, 216, 213, 213),
                  ),
                ),
                const SizedBox(
                  height: 10.0,
                ),
                const Text(
                  'category',
                  textAlign: TextAlign.start,
                ),
                const SizedBox(
                  height: 10.0,
                ),
                TextField(
                  controller: typeController,
                  decoration: InputDecoration(
                    border: OutlineInputBorder(
                        borderSide: BorderSide.none,
                        borderRadius: BorderRadius.circular(10)),
                    filled: true,
                    fillColor: const Color.fromARGB(169, 216, 213, 213),
                  ),
                ),
                const SizedBox(
                  height: 10.0,
                ),
                const Text(
                  'price',
                  textAlign: TextAlign.start,
                ),
                const SizedBox(
                  height: 10.0,
                ),
                TextField(
                  controller: priceController,
                  decoration: InputDecoration(
                    border: OutlineInputBorder(
                        borderSide: BorderSide.none,
                        borderRadius: BorderRadius.circular(10)),
                    filled: true,
                    fillColor: const Color.fromARGB(169, 216, 213, 213),
                  ),
                ),
                const SizedBox(
                  height: 10.0,
                ),
                const Text(
                  'description',
                  textAlign: TextAlign.start,
                ),
                const SizedBox(
                  height: 10.0,
                ),
                TextField(
                  controller: descriptionController,
                  maxLines: 5,
                  decoration: InputDecoration(
                    border: OutlineInputBorder(
                        borderSide: BorderSide.none,
                        borderRadius: BorderRadius.circular(10)),
                    filled: true,
                    fillColor: const Color.fromARGB(169, 216, 213, 213),
                  ),
                ),
                const SizedBox(
                  height: 20.0,
                ),
                BlocBuilder<ProductBloc, ProductState>(
                  builder: (context, state) {
                    return CustomButton(
                      backgroundColor: Theme.of(context).primaryColor,
                      foregroundColor: Colors.white,
                      borderColor: Theme.of(context).primaryColor,
                      buttonWidth: double.maxFinite,
                      buttonHeight: 45,
                      child: state is ProductLoading
                          ? const CircularProgressIndicator(
                              valueColor:
                                  AlwaysStoppedAnimation<Color>(Colors.white),
                            )
                          : const Text(
                              'UPDATE',
                              style: TextStyle(fontWeight: FontWeight.w600),
                            ),
                      onPressed: () {
                        
                        // Create new product
                       context.read<ProductBloc>().add(
                                  UpdateProductEvent(
                                    product: Product(
                                      id: widget.product!
                                          .id, 
                                      name: nameController.text,
                                      price: double.parse(priceController.text),
                                      description: descriptionController.text,
                                      imageUrl:
                                          '', // Handle the image URL accordingly
                                    ),
                                  ),
                                );
                      },
                    );
                  },
                ),
                const SizedBox(
                  height: 20.0,
                ),
                CustomButton(
                    backgroundColor: Colors.white,
                    foregroundColor: const Color.fromRGBO(255, 19, 19, 0.79),
                    borderColor: const Color.fromRGBO(255, 19, 19, 0.79),
                    buttonWidth: double.maxFinite,
                    buttonHeight: 45,
                    child: const Text(
                      'DELETE',
                      style: TextStyle(fontWeight: FontWeight.w600),
                    ),
                    onPressed: () {}),
              ],
            ),
          ),
        ),
      ),
    );
  }
}