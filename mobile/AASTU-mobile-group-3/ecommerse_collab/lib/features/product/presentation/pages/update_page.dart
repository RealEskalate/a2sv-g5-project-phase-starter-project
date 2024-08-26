
import 'package:flutter/cupertino.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:image_picker/image_picker.dart';
import 'package:flutter/material.dart';
import 'package:flutter/widgets.dart';

import '../../../../service_locator.dart';
import '../../../authentication/domain/entity/user.dart';
import '../../domain/entity/product.dart';
import '../../domain/usecase/add_product.dart';
import '../../domain/usecase/delete_product.dart';
import '../../domain/usecase/get_all_product.dart';
import '../../domain/usecase/get_product.dart';
import '../../domain/usecase/update_product.dart';
import '../bloc/blocs.dart';
import '../bloc/events.dart';
import '../bloc/states.dart';
import 'detail_page.dart';
import '../widgets/widgets.dart';

class UpdateProduct extends StatefulWidget {
  const UpdateProduct({super.key, required this.product, required this.user});
  final Product product;
  final User user;
  @override
  State<UpdateProduct> createState() => _UpdateProductState();
}

class _UpdateProductState extends State<UpdateProduct> {
  final ImagePicker _picker = ImagePicker();
  XFile? _image;
  String? name;
  String? category;
  double? price;
  String? description;

  Future<void> _pickImage(ImageSource source) async {
    final XFile? image = await _picker.pickImage(source: source);
    setState(() {
      _image = image;
    });
  }

  Widget banner() {
    return GestureDetector(
      onTap: () {
        // setState(() {
        //   _pickImage(ImageSource.gallery);
        //   debugPrint(_image.toString());
        // });
      },
      child: Container(
        width: 366,
        height: 200,
        decoration: BoxDecoration(
            color: const Color(0xFFF3F3F3),
            borderRadius: BorderRadius.circular(16)),
        child: Expanded(
          child: Column(
            mainAxisAlignment: _image == null
                ? MainAxisAlignment.spaceEvenly
                : MainAxisAlignment.start,
            crossAxisAlignment: CrossAxisAlignment.center,
            children: [
              // Stack(
              //   children: [
              //     // _image != null? Image.file(File(_image!.path)):
              //     SizedBox(
              //         width: 366,
              //         height: 200,
              //         child: Image.network(
              //           widget.product.image,
              //           fit: BoxFit.cover,
              //         )),
              //     const Positioned(
              //       child: Icon(
              //         Icons.update,
              //         size: 40,
              //       ),
              //     ),
              //   ],
              // )
              // , const Text(
              //       "update image",
              //       style: TextStyle(
              //           fontWeight: FontWeight.w500, color: Color(0xFF3E3E3E)),
              //     )
            ],
          ),
        ),
      ),
    );
  }

  @override
  Widget build(BuildContext context) {
    var name = TextEditingController(text: widget.product.name);
    var description = TextEditingController(text: widget.product.description);
    var price = TextEditingController(text: widget.product.price.toString());
    return BlocProvider(
      create: (context) => ProductBloc(
                addProductUseCase: getIt<AddProductUseCase>(),
                deleteProductUseCase: getIt<DeleteProductUseCase>(),
                updateProductUseCase: getIt<UpdateProductUseCase>(),
                getAllProductUseCase: getIt<GetAllProductUseCase>(),
                getProductUseCase: getIt<GetProductUseCase>()),
      child: Scaffold(
        appBar: AppBar(
          title:
              const Align(alignment: Alignment.center, child: Text("Update")),
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
                    TextFieldCard(
                      "name",
                      text: name,
                    ),
                    TextFieldCard(
                      "category",
                      text: TextEditingController(
                          text: widget.product.category ?? ""),
                    ),
                    TextFieldCard("price", dollar: true, text: price),
                    TextFieldCard(
                      "description",
                      area: true,
                      text: description,
                    ),
                    Container(
                      margin: const EdgeInsets.all(6),
                      child: ElevatedButton(
                        style: ElevatedButton.styleFrom(
                            backgroundColor:  const Color(0xFF3E50F3),
                            minimumSize: const Size(366, 50),
                            shape: RoundedRectangleBorder(
                                borderRadius: BorderRadius.circular(8))),
                        onPressed: () {
                          context.read<ProductBloc>().add(
                                UpdateProductEvent(
                                  productId: widget.product.id,
                                  newName: name.text,
                                  newDescription: description.text,
                                  newPrice: double.parse(price.text),
                                ),
                              );

                          if (state is SuccessState) {
                            print('');
                            ScaffoldMessenger.of(context).showSnackBar(
                              SnackBar(
                                content: Text(state.message),
                                backgroundColor: Colors.green,
                              ),
                            );
                          } else if (state is ErrorState) {
                            ScaffoldMessenger.of(context).showSnackBar(
                              SnackBar(
                                content: Text(state.message),
                                backgroundColor: Colors.red,
                              ),
                            );
                          }
                            //delay
                          Future.delayed(const Duration(seconds: 2), () { 
                            Navigator.of(context).push(MaterialPageRoute(
                              builder: (BuildContext context) {
                              final product = Product(
                                id: widget.product.id,
                                name: name.text,
                                category: category,
                                price: double.parse(price.text),
                                description: description.text,
                                image: widget.product.image,seller: widget.product.seller
                              );
                            return DetailPage(product: product, user: widget.user);
                          }));
                          });
                          Navigator.of(context).push(MaterialPageRoute(
                              builder: (BuildContext context) {
                              final product = Product(
                                id: widget.product.id,
                                name: name.text,
                                category: category,
                                price: double.parse(price.text),
                                description: description.text,
                                image: widget.product.image,
                                seller: widget.product.seller
                              );
                            return DetailPage(product: product, user: widget.user,);
                          }));
                        },
                        child: const Text(
                          "Update",
                          style: TextStyle(color: Colors.white),
                        ),
                      ),
                    ),
                    Container(
                      margin: const EdgeInsets.all(6),
                      child: OutlinedButton(
                          style: ButtonStyle(
                              minimumSize: WidgetStateProperty.all<Size>(
                                  const Size(366, 50)),
                              side: WidgetStateProperty.all<BorderSide>(
                                  const BorderSide(
                                color: Colors.red,
                                width: 1,
                              )),
                              shape: WidgetStateProperty.all<
                                      RoundedRectangleBorder>(
                                  RoundedRectangleBorder(
                                      borderRadius:
                                          BorderRadius.circular(12)))),
                          onPressed: () {
                            // Navigator.of(context).push(MaterialPageRoute(
                            //     builder: (BuildContext context) {
                            //   return const HomePage();
                            // }));
                          },
                          child: const Text("DELETE",
                              style: TextStyle(color: Colors.red))),
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
