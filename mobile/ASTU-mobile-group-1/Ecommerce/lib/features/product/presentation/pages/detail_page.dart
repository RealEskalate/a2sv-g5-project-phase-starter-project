import 'dart:developer';

import '../../../../config/route/route.dart' as route;
import 'pages.dart';

class DetailPage extends StatefulWidget {
  const DetailPage({super.key});
  @override
  State<DetailPage> createState() => _DetailPageState();
}

class _DetailPageState extends State<DetailPage> {
  late int currentSize;

  @override
  void initState() {
    currentSize = 2;
    super.initState();
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: SafeArea(
        child: BlocListener<ProductBloc, ProductState>(
          listener: (context, state) {
            log('detail + ${state}');
            if (state is ShowMessageState) {
              showCustomSnackBar(context, state.message);
              context.read<ProductBloc>().add(ResetMessageStateEvent());
            } else if (state is ErrorState) {
              showCustomSnackBar(context, state.message);
              context.read<ProductBloc>().add(LoadAllProductEvent());
            } else if (state is LoadedAllProductsState) {
              Navigator.pushReplacementNamed(context, route.homePage);
            }
          },
          child: BlocBuilder<ProductBloc, ProductState>(
            builder: (context, state) {
              if (state is LoadedSingleProductState) {
                return Column(
                  children: [
                    Stack(children: [
                      imageLoader(state.product.imageUrl),
                      BackButtonWidget.backButtonWidget(
                        iconColor: Colors.indigo,
                        onTap: () {
                          Navigator.of(context).pop();
                          context
                              .read<ProductBloc>()
                              .add(LoadAllProductEvent());
                        },
                      )
                    ]),
                    Padding(
                      padding: const EdgeInsets.only(
                        left: 16,
                        right: 16,
                      ),
                      child: Column(
                        children: [
                          const SizedBox(
                            height: 10,
                          ),
                          const Row(
                            children: [
                              CustomText(
                                text: 'Category',
                                color: Colors.grey,
                                fontSize: 16,
                              ),
                              Spacer(),
                              Icon(
                                Icons.star,
                                color: Colors.yellow,
                              ),
                              CustomText(
                                text: '(4.0)',
                                color: Colors.grey,
                                fontSize: 16,
                              )
                            ],
                          ),
                          Row(
                            children: [
                              Text(
                                state.product.name,
                                style: const TextStyle(
                                  fontSize: 24,
                                  fontWeight: FontWeight.w500,
                                ),
                              ),
                              const Spacer(),
                              Text(
                                '\$${state.product.price}',
                                style: const TextStyle(
                                  fontSize: 14,
                                  fontWeight: FontWeight.w500,
                                ),
                              ),
                            ],
                          ),
                          const SizedBox(
                            height: 20,
                          ),
                          const Align(
                            alignment: Alignment.topLeft,
                            child: CustomText(
                              text: 'Size:',
                              fontSize: 20,
                              fontWeight: FontWeight.w500,
                            ),
                          ),
                          SizedBox(
                            height: 60,
                            child: SingleChildScrollView(
                              scrollDirection: Axis.horizontal,
                              child: Row(
                                children: List.generate(
                                  6,
                                  (idx) => GestureDetector(
                                    onTap: () {
                                      setState(() {
                                        currentSize = idx;
                                      });
                                    },
                                    child: Container(
                                      width: 60,
                                      height: 60,
                                      decoration: BoxDecoration(
                                          color: idx == currentSize
                                              ? const Color(0xFF3F51F3)
                                              : null,
                                          borderRadius: const BorderRadius.all(
                                            Radius.circular(8),
                                          )),
                                      child: Center(
                                          child: CustomText(
                                        text: '${idx + 39}',
                                        fontSize: 20,
                                        color: idx == currentSize
                                            ? Colors.white
                                            : Colors.black,
                                        fontWeight: FontWeight.bold,
                                      )),
                                    ),
                                  ),
                                ),
                              ),
                            ),
                          ),
                          const SizedBox(
                            height: 10,
                          ),
                          Padding(
                            padding: const EdgeInsets.all(16),
                            child: Align(
                              alignment: Alignment.topLeft,
                              child: CustomText(
                                text: state.product.description,
                                fontWeight: FontWeight.w500,
                                color: Colors.grey,
                              ),
                            ),
                          ),
                          Row(
                            children: [
                              CustomOutlinedButton(
                                text: 'DELETE',
                                width: 150,
                                height: 50,
                                color: Colors.red,
                                onPressed: () {
                                  context.read<ProductBloc>().add(
                                      DeleteProductEvent(id: state.product.id));
                                },
                              ),
                              const Spacer(),
                              CustomOutlinedButton(
                                text: 'UPDATE',
                                width: 150,
                                height: 50,
                                backgroundColor: const Color(0xFF3F51F3),
                                color: Colors.white,
                                onPressed: () {
                                  Navigator.pushNamed(
                                      context, route.addUpdatePage,
                                      arguments: {
                                        'isUpdate': true,
                                        'product': state.product,
                                      });
                                },
                              )
                            ],
                          )
                        ],
                      ),
                    ),
                  ],
                );
              } else {
                return const Center(
                  child: CircularProgressIndicator(),
                );
              }
            },
          ),
        ),
      ),
    );
  }
}


// return Column(
//                   // mainAxisAlignment: MainAxisAlignment.center,
//                   children: [
//                     Align(
//                       alignment: Alignment.topLeft,
//                       child: BackButtonWidget(
//                         iconColor: Colors.black,
//                         onTap: () {
//                           Navigator.of(context).pop();
//                           context
//                               .read<ProductBloc>()
//                               .add(LoadAllProductEvent());
//                         },
//                       ),
//                     ),
//                     SizedBox(
//                       height: MediaQuery.sizeOf(context).height * 0.4,
//                     ),
//                     const Text(
//                       'Failed to load Product',
//                       style: TextStyle(
//                           color: Colors.red,
//                           fontSize: 20,
//                           fontWeight: FontWeight.bold),
//                     ),
//                   ],
//                 );