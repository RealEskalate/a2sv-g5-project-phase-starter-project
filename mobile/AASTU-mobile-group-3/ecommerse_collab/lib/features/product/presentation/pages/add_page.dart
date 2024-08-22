import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:image_picker/image_picker.dart';

import '../../../../service_locator.dart';
import '../../domain/entity/product.dart';
import '../../domain/usecase/add_product.dart';
import '../../domain/usecase/delete_product.dart';
import '../../domain/usecase/get_all_product.dart';
import '../../domain/usecase/get_product.dart';
import '../../domain/usecase/update_product.dart';
import '../bloc/blocs.dart';
import '../bloc/events.dart';
import '../bloc/states.dart';
import '../widgets/widgets.dart';

class AddProduct extends StatefulWidget {
  const AddProduct({super.key});

  @override
  State<AddProduct> createState() => _AddProductState();
}

class _AddProductState extends State<AddProduct> {
  final ImagePicker _picker = ImagePicker();
  XFile? _image;
  TextEditingController name = TextEditingController();
  TextEditingController category = TextEditingController();
  TextEditingController price = TextEditingController();
  TextEditingController description = TextEditingController();

  Future<void> _pickImage(ImageSource source) async {
    final XFile? image = await _picker.pickImage(source: source);
    setState(() {
      _image = image;
    });
  }

  Widget banner() {
    return GestureDetector(
      onTap: () {
        setState(() {
          _pickImage(ImageSource.gallery);
          debugPrint("Imagachen ${_image.toString()}");
        });
      },
      child: Container(
        width: 366,
        height: 160,
        padding: const EdgeInsets.all(50),
        decoration: BoxDecoration(
          color: const Color(0xFFF3F3F3),
          borderRadius: BorderRadius.circular(16),
        ),
        child: Column(
          mainAxisAlignment: _image == null
              ? MainAxisAlignment.spaceEvenly
              : MainAxisAlignment.start,
          crossAxisAlignment: CrossAxisAlignment.center,
          children: [
            const Icon(
              Icons.image_outlined,
              size: 40,
            ),
            (_image == null)
                ? const Text(
                    "upload image",
                    style: TextStyle(
                        fontWeight: FontWeight.w500, color: Color(0xFF3E3E3E)),
                  )
                : const Text(
                    "Image Uploaded",
                    style: TextStyle(
                        fontWeight: FontWeight.w500, color: Color(0xFF3E3E3E)),
                  )
          ],
        ),
      ),
    );
  }

  @override
  Widget build(BuildContext context) {
    return BlocProvider(
      create: (context) => ProductBloc(
        addProductUseCase: getIt<AddProductUseCase>(),
        deleteProductUseCase: getIt<DeleteProductUseCase>(),
        updateProductUseCase: getIt<UpdateProductUseCase>(),
        getAllProductUseCase: getIt<GetAllProductUseCase>(),
        getProductUseCase: getIt<GetProductUseCase>(),
      ),
      child: Scaffold(
        appBar: AppBar(
          title: const Align(
            alignment: Alignment.center,
            child: Text("Add Product"),
          ),
          leading: IconButton(
            icon: const Icon(Icons.arrow_back_ios),
            onPressed: () {
              Navigator.of(context).pop();
            },
          ),
        ),
        body: SingleChildScrollView(
          child: Padding(
            padding: const EdgeInsets.symmetric(horizontal: 10.0),
            child: BlocBuilder<ProductBloc, ProductState>(
              builder: (context, state) {
                return Column(
                  mainAxisAlignment: MainAxisAlignment.spaceEvenly,
                  children: [
                    banner(),
                    TextFieldCard("name", text: name),
                    TextFieldCard("category", text: category),
                    TextFieldCard("price", dollar: true, text: price),
                    TextFieldCard("description", area: true, text: description),
                    Container(
                      margin: const EdgeInsets.all(6),
                      child: ElevatedButton(
                        style: ElevatedButton.styleFrom(
                          backgroundColor: Colors.blue[800],
                          minimumSize: const Size(366, 50),
                          shape: RoundedRectangleBorder(
                            borderRadius: BorderRadius.circular(8),
                          ),
                        ),
                        onPressed: () {
                          context.read<ProductBloc>().add(
                            AddProductEvent(
                              product: Product(
                                id: '',
                                name: name.text,
                                category: category.text,
                                price: int.parse(price.text),
                                description: description.text,
                                image: _image?.path ?? "",
                              ),
                            ),
                          );
                          print(state);
                          if (state is InitialState) {
                            print('loading Mihret');
                          } else if (state is AddProductState) {
                            print("Kiki ${state.product}");
                            SnackBar(
                              content: Text(state.product.name),
                            );
                          } else if (state is ErrorState) {
                            SnackBar(
                              content: Text(state.message),
                            );
                          } else {
                            print("Error");
                          }
                          // delay
                          Future.delayed(const Duration(seconds: 2), () {
                            Navigator.of(context).pop();
                          });
                        },
                        child: const Text(
                          "ADD",
                          style: TextStyle(
                            color: Color.fromARGB(255, 210, 206, 206),
                          ),
                        ),
                      ),
                    ),
                    Container(
                      margin: const EdgeInsets.all(6),
                      child: BlocBuilder<ProductBloc, ProductState>(
                        builder: (context, state) {
                          return OutlinedButton(
                            style: ButtonStyle(
                              minimumSize: WidgetStateProperty.all<Size>(
                                const Size(366, 50),
                              ),
                              side: WidgetStateProperty.all<BorderSide>(
                                const BorderSide(
                                  color: Colors.red,
                                  width: 1,
                                ),
                              ),
                              shape: WidgetStateProperty.all<
                                  RoundedRectangleBorder>(
                                RoundedRectangleBorder(
                                  borderRadius: BorderRadius.circular(12),
                                ),
                              ),
                            ),
                            onPressed: () {
                              Navigator.of(context).pop();
                            },
                            child: const Text(
                              "DELETE",
                              style: TextStyle(color: Colors.red),
                            ),
                          );
                        },
                      ),
                    ),
                  ],
                );
              },
            ),
          ),
        ),
      ),
    );
  }
}
