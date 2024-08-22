import 'dart:io';

import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:image_picker/image_picker.dart';

import '../../../../injection_container.dart';
import '../../domain/entitity/product.dart';
import '../bloc/product_bloc.dart';
import '../bloc/product_event.dart';
import '../bloc/product_state.dart';
import '../widgets/app_bar.dart';
import '../widgets/text_field.dart';

class AddProductPage extends StatefulWidget {
  const AddProductPage({super.key});

  @override
  State<AddProductPage> createState() => _AddProductPageState();
}

class _AddProductPageState extends State<AddProductPage> {
  String? filePath;

  Future<void> pickImage() async {
    final ImagePicker picker = ImagePicker();
    final XFile? image = await picker.pickImage(source: ImageSource.gallery);

    if (image != null) {
      setState(() {
        filePath = image.path;
      });
    }
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: const MyAppBar(
        title: 'Add Product',
      ),
      body: buildBody(context),
    );
  }

  Widget buildBody(BuildContext context) {
    final Map arguments =
        (ModalRoute.of(context)?.settings.arguments as Map?) ?? {};
    final Product? product = arguments['product'];

    final TextEditingController nameController =
        TextEditingController(text: product?.name ?? '');
    final TextEditingController categoryController =
        TextEditingController(); // Separate controller for category
    final TextEditingController priceController =
        TextEditingController(text: product?.price.toString() ?? '');
    final TextEditingController descriptionController =
        TextEditingController(text: product?.description ?? '');
    String? imageUrl = product?.imageUrl;

    final ProductBloc productBloc = BlocProvider.of<ProductBloc>(context);

    return SingleChildScrollView(
      padding: const EdgeInsets.all(20),
      child: Column(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          GestureDetector(
            onTap: () async {
              await pickImage();
            },
            child: filePath != null || imageUrl != null
                ? AspectRatio(
                    aspectRatio: 16 / 7,
                    child: filePath != null
                        ? Image.file(
                            File(filePath!),
                            fit: BoxFit.fill,
                          )
                        : Image.network(
                            imageUrl!,
                            fit: BoxFit.fill,
                          ),
                  )
                : Container(
                    padding: const EdgeInsets.symmetric(vertical: 65),
                    decoration: BoxDecoration(
                      borderRadius: BorderRadius.circular(10),
                      color: const Color.fromARGB(155, 232, 229, 229),
                    ),
                    child: const Center(
                      child: Column(
                        children: [
                          Icon(
                            Icons.image,
                            size: 50,
                          ),
                          SizedBox(
                            height: 20,
                          ),
                          Text('Upload Image'),
                        ],
                      ),
                    ),
                  ),
          ),
          const SizedBox(
            height: 20,
          ),
          MyTextField(
            controller: nameController,
            lable: 'Name',
            lines: 1,
          ),
          MyTextField(
            controller: categoryController,
            lable: 'Category',
            lines: 1,
          ),
          MyTextField(
            controller: priceController,
            lable: 'Price',
            lines: 1,
            suffIcon: const Icon(
              Icons.attach_money,
              color: Colors.black,
            ),
          ),
          MyTextField(
            controller: descriptionController,
            lable: 'Description',
            lines: 5,
          ),
          const SizedBox(
            height: 50,
          ),
          Column(
            mainAxisAlignment: MainAxisAlignment.end,
            crossAxisAlignment: CrossAxisAlignment.center,
            children: [
              ElevatedButton(
                onPressed: () {
                  final newProduct = Product(
                    id: product?.id ?? '',
                    name: nameController.text,
                    price: double.parse(priceController.text),
                    description: descriptionController.text,
                    imageUrl: filePath ?? imageUrl ?? '',
                  );

                  if (product == null) {
                    productBloc.add(InsertProductEvent(product: newProduct));
                    Navigator.pushNamed(context, '/homepage');
                  } else {
                    productBloc.add(UpdateProductEvent(product: newProduct));
                    Navigator.pushNamed(context, '/homepage');
                  }
                },
                style: ButtonStyle(
                  backgroundColor: MaterialStateProperty.all<Color>(
                    const Color.fromARGB(255, 32, 77, 202),
                  ),
                  shape: MaterialStateProperty.all<RoundedRectangleBorder>(
                    RoundedRectangleBorder(
                      borderRadius: BorderRadius.circular(10),
                    ),
                  ),
                ),
                child: Container(
                  padding: const EdgeInsets.all(15),
                  width: double.infinity,
                  child: const Text(
                    'ADD',
                    style: TextStyle(color: Colors.white),
                    textAlign: TextAlign.center,
                  ),
                ),
              ),
              const SizedBox(
                height: 20,
              ),
              ElevatedButton(
                onPressed: () {
                  Navigator.pop(context);
                },
                style: ButtonStyle(
                  backgroundColor: MaterialStateProperty.all<Color>(
                    Colors.white,
                  ),
                  foregroundColor: MaterialStateProperty.all<Color>(
                    Colors.red,
                  ),
                  side: MaterialStateProperty.all<BorderSide>(
                    const BorderSide(color: Colors.red, width: 1),
                  ),
                  shape: MaterialStateProperty.all<RoundedRectangleBorder>(
                    RoundedRectangleBorder(
                      borderRadius: BorderRadius.circular(10),
                    ),
                  ),
                ),
                child: Container(
                  padding: const EdgeInsets.all(15),
                  width: double.infinity,
                  child: const Text(
                    'DELETE',
                    style: TextStyle(color: Colors.red),
                    textAlign: TextAlign.center,
                  ),
                ),
              ),
            ],
          ),
        ],
      ),
    );
  }
}
