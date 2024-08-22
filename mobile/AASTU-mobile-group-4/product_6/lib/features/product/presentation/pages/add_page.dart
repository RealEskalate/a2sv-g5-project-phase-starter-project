import 'dart:io';
import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:image_picker/image_picker.dart';

import '../../domain/entities/product.dart';
import '../bloc/product_bloc.dart';
import '../bloc/product_event.dart';
import '../bloc/product_state.dart';
import 'home_page.dart';

class AddPage extends StatelessWidget {
  final _formKey = GlobalKey<FormState>();

  final _nameController = TextEditingController();
  final _categoryController = TextEditingController();
  final _priceController = TextEditingController();
  final _descriptionController = TextEditingController();

  final ImagePicker _picker = ImagePicker();
  File? _image;

  AddPage({super.key});

  Future<void> _pickImage(BuildContext context) async {
    final XFile? image = await _picker.pickImage(source: ImageSource.gallery);
    if (image != null) {
      _image = File(image.path);
      (context as Element).markNeedsBuild(); // To trigger rebuild
    }
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: Colors.white,
      appBar: AppBar(
        backgroundColor: Colors.white,
        title: const Text(
          "Add Product",
          style: TextStyle(fontSize: 17, fontWeight: FontWeight.bold),
        ),
        centerTitle: true,
        automaticallyImplyLeading: false,
        leading: IconButton(
          onPressed: () {
            Navigator.pop(context);
            context.read<ProductBloc>().add(LoadAllProductEvent());
          },
          icon: const Icon(
            Icons.arrow_back_ios,
            color: Colors.blue,
          ),
        ),
      ),
      body: BlocListener<ProductBloc, ProductState>(
        listener: (context, state) {
          if (state is Success) {
            Navigator.push(
              context,
              MaterialPageRoute(
                builder: (context) {
                  return HomePage();
                },
              ),
            );
            ScaffoldMessenger.of(context).showSnackBar(
              const SnackBar(content: Text('Product added successfully!')),
            );
          } else if (state is ErrorState) {
            ScaffoldMessenger.of(context).showSnackBar(
              SnackBar(content: Text(state.message)),
            );
          }
        },
        child: SingleChildScrollView(
          child: Column(
            children: [
              GestureDetector(
                onTap: () => _pickImage(context),
                child: Container(
                  width: 366,
                  height: 190,
                  margin: const EdgeInsets.all(18),
                  decoration: BoxDecoration(
                    borderRadius: BorderRadius.circular(15),
                    color: const Color.fromARGB(255, 223, 219, 219),
                  ),
                  child: Center(
                    child: _image == null
                        ? Column(
                            mainAxisAlignment: MainAxisAlignment.center,
                            crossAxisAlignment: CrossAxisAlignment.center,
                            children: [
                              Icon(Icons.image,
                                  color: Colors.grey[600], size: 50),
                              const SizedBox(height: 10),
                              const Text(
                                "Upload image",
                                style: TextStyle(fontWeight: FontWeight.bold),
                              ),
                            ],
                          )
                        : Image.file(
                            _image!,
                            height: 200,
                            width: double.infinity,
                            fit: BoxFit.cover,
                          ),
                  ),
                ),
              ),
              Padding(
                padding: const EdgeInsets.all(23.0),
                child: Form(
                  key: _formKey,
                  child: Column(
                    crossAxisAlignment: CrossAxisAlignment.start,
                    children: <Widget>[
                      const Text(
                        'Name',
                        style: TextStyle(fontSize: 16),
                      ),
                      TextFormField(
                        controller: _nameController,
                        decoration: const InputDecoration(
                          filled: true,
                          fillColor: Color.fromARGB(255, 223, 219, 219),
                          border: OutlineInputBorder(
                            borderSide: BorderSide.none,
                          ),
                          contentPadding: EdgeInsets.all(8),
                        ),
                        validator: (value) {
                          if (value == null || value.isEmpty) {
                            return 'Please enter a product name';
                          }
                          return null;
                        },
                      ),
                      const SizedBox(height: 16),
                      const Text(
                        'Category',
                        style: TextStyle(fontSize: 16),
                      ),
                      TextFormField(
                        controller: _categoryController,
                        decoration: const InputDecoration(
                          filled: true,
                          fillColor: Color.fromARGB(255, 223, 219, 219),
                          border: OutlineInputBorder(
                            borderSide: BorderSide.none,
                          ),
                          contentPadding: EdgeInsets.all(8),
                        ),
                        validator: (value) {
                          if (value == null || value.isEmpty) {
                            return 'Please enter a product category';
                          }
                          return null;
                        },
                      ),
                      const SizedBox(height: 16),
                      const Text(
                        'Price',
                        style: TextStyle(fontSize: 16),
                      ),
                      TextFormField(
                        controller: _priceController,
                        decoration: const InputDecoration(
                          filled: true,
                          fillColor: Color.fromARGB(255, 223, 219, 219),
                          border: OutlineInputBorder(
                            borderSide: BorderSide.none,
                          ),
                          contentPadding: EdgeInsets.all(8),
                          suffixIcon: Icon(Icons.attach_money),
                        ),
                        keyboardType: TextInputType.number,
                        validator: (value) {
                          if (value == null || value.isEmpty) {
                            return 'Please enter a price';
                          }
                          return null;
                        },
                      ),
                      const SizedBox(height: 16),
                      const Text(
                        'Description',
                        style: TextStyle(fontSize: 16),
                      ),
                      TextFormField(
                        controller: _descriptionController,
                        decoration: const InputDecoration(
                          filled: true,
                          fillColor: Color.fromARGB(255, 223, 219, 219),
                          border:
                              OutlineInputBorder(borderSide: BorderSide.none),
                          contentPadding: EdgeInsets.all(8),
                        ),
                        maxLines: 4,
                        validator: (value) {
                          if (value == null || value.isEmpty) {
                            return 'Please enter a description';
                          }
                          return null;
                        },
                      ),
                      const SizedBox(height: 20),
                      ElevatedButton(
                        onPressed: () {
                          if (_formKey.currentState!.validate()) {
                            final product = Product(
                              id: '', // Assuming ID is generated by backend
                              name: _nameController.text,
                              description: _descriptionController.text,
                              price: double.parse(_priceController.text),
                              imageUrl: _image != null ? _image!.path : '',
                            );

                            context
                                .read<ProductBloc>()
                                .add(CreateProductEvent(product: product));
                          }
                        },
                        style: ElevatedButton.styleFrom(
                          minimumSize: const Size(500, 50),
                          padding: const EdgeInsets.all(16),
                          backgroundColor: Colors.blue,
                          shape: RoundedRectangleBorder(
                            borderRadius: BorderRadius.circular(8),
                          ),
                        ),
                        child: const Text(
                          'ADD',
                          style: TextStyle(color: Colors.white),
                        ),
                      ),
                      const SizedBox(height: 10),
                      OutlinedButton(
                        onPressed: () {
                          // Logic for delete functionality
                        },
                        style: OutlinedButton.styleFrom(
                          minimumSize: const Size(500, 50),
                          padding: const EdgeInsets.all(16),
                          side: const BorderSide(color: Colors.red),
                          shape: RoundedRectangleBorder(
                            borderRadius: BorderRadius.circular(8),
                          ),
                        ),
                        child: const Text(
                          "DELETE",
                          style: TextStyle(color: Colors.red),
                        ),
                      ),
                    ],
                  ),
                ),
              ),
            ],
          ),
        ),
      ),
    );
  }
}






