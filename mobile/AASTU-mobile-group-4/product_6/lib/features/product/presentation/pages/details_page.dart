import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:google_fonts/google_fonts.dart';

import '../../domain/entities/product.dart';
import '../bloc/product_bloc.dart';
import '../bloc/product_event.dart';
import '../bloc/product_state.dart';
import 'home_page.dart';
import 'update_page.dart';

class DetailsPage extends StatefulWidget {
  final Product productObject;
  const DetailsPage({super.key, required this.productObject});

  @override
  State<DetailsPage> createState() => _DetailsPageState();
}

class _DetailsPageState extends State<DetailsPage> {
  int? selectedIndex;
  @override
  Widget build(BuildContext context) {
    return SafeArea(
      child: Scaffold(
        body: ListView(
          children: [
            Stack(
              children: [
                Container(
                  color: const Color.fromRGBO(190, 162, 155, 350),
                  child: Image.network(widget.productObject.imageUrl),
                ),
                Positioned(
                  right: 304,
                  bottom: 104,
                  child: GestureDetector(
                    onTap: () {
                      Navigator.pop(context);
                      context.read<ProductBloc>().add(LoadAllProductEvent());
                      // Navigator.push(
                      //   context,
                      //   MaterialPageRoute(
                      //     builder: (context) => HomePage(),
                      //   ),
                      // );
                    },
                    child: Container(
                      height: 40,
                      width: 40,
                      decoration: const BoxDecoration(
                          color: Colors.white, shape: BoxShape.circle),
                      child: const Center(
                        child: Icon(
                          Icons.arrow_back_ios,
                          size: 18,
                          color: Color.fromRGBO(63, 81, 243, 1),
                        ),
                      ),
                    ),
                  ),
                ),
              ],
            ),
            const SizedBox(height: 10),
            Padding(
              padding: const EdgeInsets.only(left: 32.0, right: 32.0),
              child: Row(
                mainAxisAlignment: MainAxisAlignment.spaceBetween,
                children: [
                  Text(
                    'men',
                    style: GoogleFonts.poppins(
                      textStyle: const TextStyle(
                        fontWeight: FontWeight.w400,
                        color: Color.fromRGBO(170, 170, 170, 1),
                        fontSize: 16.0,
                        height: 0.5,
                      ),
                    ),
                  ),
                  Row(
                    // mainAxisAlignment: MainAxisAlignment.end,
                    children: [
                      Container(
                        padding: const EdgeInsets.all(0),
                        height: 30,
                        width: 30,
                        child: Center(
                          child: IconButton(
                            onPressed: () {},
                            icon: const Icon(
                              Icons.star,
                              size: 16.0,
                              color: Color.fromRGBO(255, 215, 0, 1),
                            ),
                          ),
                        ),
                      ),
                      Text(
                        '(4.0)',
                        style: GoogleFonts.sora(
                            textStyle: const TextStyle(
                          fontWeight: FontWeight.w400,
                          fontSize: 16.0,
                          color: Color.fromRGBO(170, 170, 170, 1),
                        )),
                      ),
                    ],
                  )
                ],
              ),
            ),
            Padding(
              padding: const EdgeInsets.only(left: 32.0, right: 32.0),
              child: Row(
                mainAxisAlignment: MainAxisAlignment.spaceBetween,
                children: [
                  Text(
                    widget.productObject.name,
                    style: GoogleFonts.poppins(
                      textStyle: const TextStyle(
                        fontWeight: FontWeight.w600,
                        fontSize: 24,
                        color: Color.fromRGBO(62, 62, 62, 1),
                      ),
                    ),
                  ),
                  Text(
                    widget.productObject.price.toString(),
                    style: GoogleFonts.poppins(
                        textStyle: const TextStyle(
                      fontWeight: FontWeight.w500,
                      fontSize: 16.0,
                      color: Color.fromRGBO(62, 62, 62, 1),
                    )),
                  )
                ],
              ),
            ),
            const SizedBox(height: 20),
            Row(
              children: [
                Padding(
                  padding: const EdgeInsets.only(left: 32.0, right: 32.0),
                  child: Text(
                    'Size:',
                    style: GoogleFonts.poppins(
                        textStyle: const TextStyle(
                      fontWeight: FontWeight.w500,
                      fontSize: 20,
                      color: Color.fromRGBO(62, 62, 62, 1),
                    )),
                  ),
                ),
              ],
            ),
            const SizedBox(height: 8),
            Container(
              padding: const EdgeInsets.symmetric(horizontal: 32),
              height: 60,
              child: ListView.builder(
                scrollDirection: Axis.horizontal,
                itemCount: 10,
                itemBuilder: (context, index) {
                  int displayIndex = 39 + index;

                  Color sizeColor = selectedIndex == index
                      ? const Color.fromRGBO(63, 81, 243, 1)
                      : Colors.white;
                  Color textColor =
                      selectedIndex == index ? Colors.white : Colors.black;

                  return GestureDetector(
                    onTap: () {
                      setState(() {
                        selectedIndex = index;
                      });
                    },
                    child: Container(
                      width: 60,
                      margin: const EdgeInsets.symmetric(horizontal: 5),
                      decoration: BoxDecoration(
                        borderRadius: BorderRadius.circular(13),
                        color: sizeColor,
                        shape: BoxShape.rectangle,
                        // boxShadow: [
                        //   BoxShadow(
                        //     color: Colors.grey.withOpacity(0.5),
                        //     spreadRadius: 2,
                        //     blurRadius: 3,
                        //     offset: const Offset(0, 5),
                        //   ),
                        // ],
                      ),
                      child: Center(
                        child: Text(
                          '$displayIndex',
                          style: TextStyle(
                            color: textColor,
                            fontSize: 20,
                          ),
                        ),
                      ),
                    ),
                  );
                },
              ),
            ),
            const SizedBox(height: 14),
            Padding(
              padding: const EdgeInsets.symmetric(horizontal: 32),
              child: Text(
                widget.productObject.description,
                style: GoogleFonts.poppins(
                    textStyle: const TextStyle(
                  fontWeight: FontWeight.w500,
                  fontSize: 14,
                  color: Color.fromRGBO(102, 102, 102, 1),
                )),
              ),
            ),
            const SizedBox(height: 8),
            _buttons(context, widget.productObject),
          ],
        ),
      ),
    );
  }
}

Widget _horizontal_scrolling(BuildContext context) {
  return Container(
    padding: const EdgeInsets.symmetric(horizontal: 32),
    height: 60,
    child: ListView.builder(
      scrollDirection: Axis.horizontal,
      itemCount: 10,
      itemBuilder: (context, index) {
        int displayIndex = 39 + index;

        Color sizeColor;
        Color textColor;

        if (index == 41) {
          sizeColor = const Color.fromRGBO(63, 81, 243, 1);
          textColor = Colors.white;
        } else {
          sizeColor = Colors.white;
          textColor = Colors.black;
        }

        return GestureDetector(
          onTap: () {
            // setState(() {
            //     selectedIndex = index;
            //   });
          },
          child: Container(
            width: 60,
            margin: const EdgeInsets.symmetric(horizontal: 5),
            decoration: BoxDecoration(
              borderRadius: BorderRadius.circular(13),
              color: sizeColor,
              shape: BoxShape.rectangle,
              // boxShadow: [
              //   BoxShadow(
              //     color: Colors.grey.withOpacity(0.5),
              //     spreadRadius: 2,
              //     blurRadius: 3,
              //     offset: const Offset(0, 5),
              //   ),
              // ],
            ),
            // color: const Color.fromRGBO(63, 81, 243, 1),
            child: Center(
              child: Text(
                '$displayIndex',
                style: const TextStyle(
                  color: Colors.black,
                  fontSize: 20,
                ),
              ),
            ),
          ),
        );
      },
    ),
  );
}

Widget _buttons(BuildContext context, Product productObject) {
  return Container(
    height: 70,
    padding: const EdgeInsets.symmetric(horizontal: 32),
    // color: Colors.black,
    child: Row(
      mainAxisAlignment: MainAxisAlignment.spaceBetween,
      children: [
        Expanded(
          child: OutlinedButton(
            style: OutlinedButton.styleFrom(
              side: const BorderSide(
                color: Colors.red,
                width: 2,
              ),
              elevation: 10,
              shape: RoundedRectangleBorder(
                borderRadius: BorderRadius.circular(12),
              ),
            ),
            onPressed: () {
              BlocProvider.of<ProductBloc>(context)
                  .add(DeleteProductEvent(productObject.id));
              BlocListener(listener: (context, state) {
                if (state is Success) {
                  Navigator.push(
                    context,
                    MaterialPageRoute(
                      builder: (context) {
                        return HomePage();
                      },
                    ),
                  );
                  context.read<ProductBloc>().add(LoadAllProductEvent());
                  ScaffoldMessenger.of(context).showSnackBar(
                    SnackBar(
                      content: Text('Successfully deleted product'),
                    ),
                  );
                } else if (state is ErrorState) {
                  ScaffoldMessenger.of(context).showSnackBar(
                    SnackBar(
                      content: Text('An error occurred'),
                    ),
                  );
                }
              });
            },
            child: Text(
              'Delete',
              style: GoogleFonts.poppins(
                color: Colors.red,
              ),
            ),
          ),
        ),
        const SizedBox(width: 50),
        Expanded(
          child: ElevatedButton(
            style: ElevatedButton.styleFrom(
              backgroundColor: const Color.fromRGBO(63, 81, 243, 1),
              elevation: 10,
              shadowColor: Colors.black,
              shape: RoundedRectangleBorder(
                borderRadius: BorderRadius.circular(12),
              ),
            ),
            onPressed: () {
              Navigator.push(
                context,
                MaterialPageRoute(
                  builder: (context) {
                    return UpdatePage(
                      product: productObject,
                    );
                  },
                ),
              );
            },
            child: Text(
              'Update',
              style: GoogleFonts.poppins(
                color: Colors.white,
              ),
            ),
          ),
        ),
      ],
    ),
  );
}
